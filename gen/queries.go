package gen

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"

	. "github.com/dave/jennifer/jen"
)

func (g *Generator) generateQueries() {
	f := NewFile(g.Package)

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

	funcz := Func().Id(id).ParamsFunc(
		func(pg *Group) {
			for _, v := range o.VariableDefinitions {
				varId := getLowerId(v.Variable)
				pg.Add(Id(varId).Add(g.getGoType(v.Type)...)) //.Comment(arg.Description))
			}
		},
	).ParamsFunc(func(pg *Group) {
		if field, ok := o.SelectionSet[0].(*ast.Field); ok {
			pg.Add(g.getGoType(field.Definition.Type)...)
		} else {
			pg.Add(Interface())
		}
		pg.Add(Error())
	}).BlockFunc(
		func(bg *Group) {
			bg.Add(Id("q").Op(":=").StructFunc(func(sg *Group) {
				g.generateSelectionSet(sg, &o.SelectionSet)
			})).Values()

			g.generateOperationVariables(bg, o)

		})

	return *funcz
}

func (g *Generator) generateSelectionSet(group *Group, set *ast.SelectionSet) {
	fmt.Println("selections")
	for _, selection := range *set {
		if field, ok := selection.(*ast.Field); ok {
			fmt.Println("\t", field.Name)
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
	}

}

func (g *Generator) generateOperationVariables(group *Group, o *ast.OperationDefinition) {
	d := Dict{}
	for _, v := range o.VariableDefinitions {
		fmt.Println("variables")
		fmt.Println("\t", v.Variable)

		d[Lit(v.Variable)] = Id(getLowerId(v.Variable))
	}
	group.Add(Id("variables").Op(":=").Map(String()).Interface().Values(d))
}
