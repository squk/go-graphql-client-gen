package main

import "github.com/squk/go-graphql-client-gen/gen"

func main() {
	g := gen.NewGenerator(
		gen.WithSchemaFile("schema.gql"),
		gen.WithQueriesFile("queries.gql"),
	)
	g.Run()

	//doc := ast.NewDocument()
	//doc.Input.ResetInputBytes(b)
	//spew.Dump(doc)

	//doc, _ := astparser.ParseGraphqlDocumentBytes(b)

	//gen := codegen.New(&doc, codegen.Config{
	//PackageName:           "types",
	//DirectiveStructSuffix: "",
	//})

	//gen.Generate(os.Stdout)
}
