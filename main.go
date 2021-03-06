package main

import (
	"github.com/squk/go-graphql-client-gen/gen"
)

func main() {
	g := gen.NewGenerator(gen.WithSchemaFile("schema.gql"))
	g.Run()
}
