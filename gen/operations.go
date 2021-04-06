package gen

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"

	. "github.com/dave/jennifer/jen"
)

func (g *Generator) generateOperations() {
	f := NewFile(g.Package)

	g.generateOperationsStruct(f)
	q := g.Queries
	for _, o := range q.Operations {
		code := g.generateOperation(o)
		f.Add(code...)
	}

	err := f.Save("types/operations.go")
	if err != nil {
		fmt.Println(err)
	}
}

func (g *Generator) generateOperationsStruct(f *File) {
	for _, name := range []string{"Queries", "Mutations"} {
		f.Type().Id(name).Struct(
			Id("client").Op("*").Qual("github.com/hasura/go-graphql-client", "Client"),
		)

		f.Func().Params(
			Id("q").Op("*").Id(name),
		).Id("SetClient").Params(
			Id("client").Op("*").Qual("github.com/hasura/go-graphql-client", "Client"),
		).Block(
			Id("q").Dot("client").Op("=").Id("client"),
		)
	}
}

func (g *Generator) generateOperation(o *ast.OperationDefinition) (code []Code) {
	operationKind := string(o.Operation)
	receiverId := string(operationKind[0])

	var clientMethod string // which method to call on go-graphql-client
	var receiverType string // which struct to make a receiver of this func
	if o.Operation == ast.Query {
		clientMethod = "Query"
		receiverType = "Queries"
	} else if o.Operation == ast.Mutation {
		clientMethod = "Mutate"
		receiverType = "Mutations"
	}

	id := strcase.ToCamel(o.Name)
	if id == "" {
		fmt.Println("unnamed operations are not supported")
		return *Comment("unnamed operations are not supported")
	}

	var opType []Code
	var selectionName string

	// we don't support multiple top level fields in query selection sets yet
	if field, ok := o.SelectionSet[0].(*ast.Field); ok {
		selectionName = strcase.ToCamel(field.Name)
		opType = g.getGoType(field.Definition.Type)
	} else {
		opType = []Code{Interface()}
	}

	funcz := Func().Params(
		Id(receiverId).Op("*").Id(receiverType),
	).Id(id).ParamsFunc( // params
		func(pg *Group) {
			for _, v := range o.VariableDefinitions {
				varId := getLowerId(v.Variable)
				if v.Definition.Kind == ast.Scalar {
					pg.Add(Id(varId).String())
				} else {
					pg.Add(Id(varId).Add(g.getGoType(v.Type)...)) //.Comment(arg.Description))
				}
			}
		},
	).ParamsFunc(func(pg *Group) { // returns
		pg.Add(opType...)
		pg.Add(Error())
	}).BlockFunc(
		func(bg *Group) {
			bg.Add(Id(operationKind).Op(":=").StructFunc(func(sg *Group) {
				g.generateSelectionSet(sg, &o.SelectionSet)
			})).Values()

			g.generateOperationVariables(bg, o)

			bg.Add(Var().Id("data").Add(opType...))
			bg.Add(Err().Op(":=").Id(receiverId).Dot("client").Dot(clientMethod).Call(
				Qual("context", "Background").Call(),
				Op("&").Id(operationKind),
				Id("variables"),
			))

			bg.Add(If().Err().Op("!=").Nil().Block(
				Return(Id("data"), Err()),
			))

			// marshal our selection set into json
			bg.Add(List(Id("bytes"), Err()).Op(":=").Qual("encoding/json", "Marshal").Call(Id(operationKind).Dot(selectionName)))
			bg.Add(If().Err().Op("!=").Nil().Block(
				Return(Id("data"), Err()),
			))

			// unmarshal selection set into generated type so it's sparsely
			// populated
			bg.Add(Err().Op("=").Qual("encoding/json", "Unmarshal").Call(Id("bytes"), Op("&").Id("data")))
			bg.Add(If().Err().Op("!=").Nil().Block(
				Return(Id("data"), Err()),
			))

			bg.Add(Return(Id("data"), Nil()))

		})

	return *funcz
}

func (g *Generator) generateSelectionSet(group *Group, set *ast.SelectionSet) {
	for _, selection := range *set {
		if field, ok := selection.(*ast.Field); ok {
			g.generateFields(group, field)
		} else if fragment, ok := selection.(*ast.FragmentSpread); ok {
			g.generateFragmentSpread(group, fragment)
		} else if fragment, ok := selection.(*ast.InlineFragment); ok {
			fmt.Printf("inline fragments not supported - %+v", fragment)
		}
	}
}

func (g *Generator) generateFragmentSpread(group *Group, fragment *ast.FragmentSpread) {
	spew.Dump(fragment.Name)
	g.generateSelectionSet(group, &fragment.Definition.SelectionSet)
}

func (g *Generator) generateFields(group *Group, field *ast.Field) {
	stmt := group.Id(strcase.ToCamel(field.Name))

	gqlId := field.Name

	// add arguments to `gql` tag
	if len(field.Arguments) != 0 {
		args := make([]string, len(field.Arguments))
		for i, a := range field.Arguments {
			args[i] = fmt.Sprintf("%s: %s", a.Name, a.Value.String())

		}
		gqlId += fmt.Sprintf("(%s)", strings.Join(args, ","))
	}

	// type is not nested, get Go type and bail
	if len(field.SelectionSet) == 0 {
		if field.Definition != nil {
			stmt.Add(g.getGoType(field.Definition.Type)...)
		}
	} else { // generate nested struct using recursion
		if typeIsArray(field.Definition.Type) {
			stmt.Index()
		}
		stmt.StructFunc(func(sg *Group) {
			g.generateSelectionSet(sg, &field.SelectionSet)
		})
	}

	tags := map[string]string{
		"graphql": gqlId,
		"json":    fmt.Sprintf("%s,%s", field.Name, "omitempty"),
	}
	stmt.Tag(tags)
}

func (g *Generator) generateOperationVariables(group *Group, o *ast.OperationDefinition) {
	d := Dict{}
	for _, v := range o.VariableDefinitions {
		variable := v.Variable
		typeName := v.Type.Name()

		if g.typeNameIsScalar(typeName) {
			if isHasuraClientScalar(typeName) {
				conversion := Qual("github.com/hasura/go-graphql-client", typeName)
				d[Lit(variable)] = conversion.Parens(Id(getLowerId(variable)))
			} else if !isGoKeyword(getLowerId(typeName)) {
				d[Lit(variable)] = Id(getScalarConstructorName(typeName)).Parens(Id(getLowerId(variable)))
			} else {
				fmt.Println("whoops", typeName)
			}
		} else {
			d[Lit(variable)] = Id(getLowerId(variable))
		}

	}
	group.Add(Id("variables").Op(":=").Map(String()).Interface().Values(d))
}

func (g *Generator) typeNameIsScalar(name string) bool {
	if t, ok := g.Schema.Types[name]; ok {
		return t.Kind == ast.Scalar
	}
	return false
}
