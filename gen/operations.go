package gen

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"

	. "github.com/dave/jennifer/jen"
)

func (g *Generator) generateQueries() {
	f := NewFile(g.Package)

	g.generateQueriesStruct(f)
	q := g.Queries
	//spew.Dump(q)
	for _, o := range q.Operations {
		code := g.generateOperation(o)
		f.Add(code...)
	}

	//fields := g.Schema.Query.Fields
	//sort.SliceStable(fields, func(i, j int) bool {
	//return strings.Compare(fields[i].Name, fields[j].Name) == -1
	//})
	err := f.Save("types/queries.go")
	if err != nil {
		fmt.Println(err)
	}
}

func (g *Generator) generateQueriesStruct(f *File) {
	f.Type().Id("Queries").Struct(
		Id("client").Op("*").Qual("github.com/hasura/go-graphql-client", "Client"),
	)

	f.Func().Params(
		Id("q").Op("*").Id("Queries"),
	).Id("SetClient").Params(
		Id("client").Op("*").Qual("github.com/hasura/go-graphql-client", "Client"),
	).Block(
		Id("q").Dot("client").Op("=").Id("client"),
	)
}

func (g *Generator) generateOperation(o *ast.OperationDefinition) (code []Code) {
	//fmt.Println(o.Name)
	//fmt.Println(o.Operation)

	id := strcase.ToCamel(o.Name)
	//for t, _ := range g.Schema.PossibleTypes {
	//spew.Dump(t)
	//}
	for _, v := range o.VariableDefinitions {
		fmt.Println("variables")
		fmt.Println("\t", v.Variable)
	}

	var opType []Code
	var selectionName string
	if field, ok := o.SelectionSet[0].(*ast.Field); ok {
		selectionName = strcase.ToCamel(field.Name)
		opType = g.getGoType(field.Definition.Type)
	} else {
		opType = []Code{Interface()}
	}

	funcz := Func().Params(
		Id("q").Op("*").Id("Queries"),
	).Id(id).ParamsFunc(
		func(pg *Group) {
			for _, v := range o.VariableDefinitions {
				varId := getLowerId(v.Variable)
				pg.Add(Id(varId).Add(g.getGoType(v.Type)...)) //.Comment(arg.Description))
			}
		},
	).ParamsFunc(func(pg *Group) {
		pg.Add(opType...)
		pg.Add(Error())
	}).BlockFunc(
		func(bg *Group) {
			bg.Add(Id("query").Op(":=").StructFunc(func(sg *Group) {
				g.generateSelectionSet(sg, &o.SelectionSet)
			})).Values()

			g.generateOperationVariables(bg, o)

			bg.Add(Err().Op(":=").Id("q").Dot("client").Dot("Query").Call(
				Qual("context", "Background").Call(),
				Op("&").Id("query"),
				Id("variables"),
			))

			bg.Add(If().Err().Op("!=").Nil().Block(
				Return(Nil(), Err()),
			))

			bg.Add(List(Id("bytes"), Err()).Op(":=").Qual("encoding/json", "Marshal").Call(Id("query").Dot(selectionName)))
			bg.Add(If().Err().Op("!=").Nil().Block(
				Return(Nil(), Err()),
			))

			bg.Add(Var().Id("data").Add(opType...))

			bg.Add(Err().Op("=").Qual("encoding/json", "Unmarshal").Call(Id("bytes"), Op("&").Id("data")))
			bg.Add(If().Err().Op("!=").Nil().Block(
				Return(Nil(), Err()),
			))

			bg.Add(Return(Id("data"), Nil()))

		})

	return *funcz
}

func (g *Generator) generateSelectionSet(group *Group, set *ast.SelectionSet) {
	for _, selection := range *set {
		if field, ok := selection.(*ast.Field); ok {
			//fmt.Println("\t", field.Name)
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
	fmt.Println(fragment.Definition)
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
		fmt.Println("variables")
		fmt.Println("\t", v.Variable)

		if g.typeNameIsScalar(v.Type.Name()) {
			d[Lit(v.Variable)] = Id(getScalarContructorName(v.Type.Name())).Parens(Id(getLowerId(v.Variable)))
		} else {
			d[Lit(v.Variable)] = Id(getLowerId(v.Variable))
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
