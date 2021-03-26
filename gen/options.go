package gen

import (
	"io/ioutil"

	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

type GeneratorOption func(*Generator)

func WithPackage(pkg string) GeneratorOption {
	return func(g *Generator) {
		g.Package = pkg
	}
}

func WithSchema(schema string) GeneratorOption {
	return func(g *Generator) {
		g.schemaSource = schema

		src := &ast.Source{
			Name:    "schema",
			Input:   schema,
			BuiltIn: false,
		}

		s, err := gqlparser.LoadSchema(src)
		if err != nil {
			panic(err)
		}
		g.Schema = s
	}
}

func WithSchemaFile(filename string) GeneratorOption {
	return func(g *Generator) {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}

		WithSchema(string(b))(g)
	}
}

func WithQueries(queries string) GeneratorOption {
	return func(g *Generator) {
		g.queriesSource = queries

		q, err := gqlparser.LoadQuery(g.Schema, queries)
		if err != nil {
			panic(err)
		}
		g.Queries = q
	}
}

func WithQueriesFile(filename string) GeneratorOption {
	return func(g *Generator) {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		WithQueries(string(b))(g)
	}
}
