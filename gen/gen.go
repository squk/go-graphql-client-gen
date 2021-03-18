package gen

import (
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"

	. "github.com/dave/jennifer/jen"
)

type Generator struct {
	Package string // package name for generated files

	Schema       *ast.Schema
	schemaSource string

	Queries       *ast.QueryDocument
	queriesSource string
}

type GeneratorOption func(*Generator)

// NewGenerator now takes a slice of option as the rest arguments
func NewGenerator(opts ...GeneratorOption) *Generator {
	const (
		defaultPackage = "types"
	)

	h := &Generator{
		Package: defaultPackage,
	}

	// Loop through each option
	for _, opt := range opts {
		// Call the option giving the instantiated
		// *Generator as the argument
		opt(h)
	}

	// return the modified house instance
	return h
}

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

func (g *Generator) Run() {
	g.generateTypes()
	g.generateQueries()
}

func (g *Generator) generateTypes() {
	enums := NewFile(g.Package)
	scalars := NewFile(g.Package)
	input := NewFile(g.Package)
	objects := NewFile(g.Package)

	defer enums.Save("types/enums.go")
	defer scalars.Save("types/scalars.go")
	defer input.Save("types/input.go")
	defer objects.Save("types/objects.go")

	// iterate over map in alphabetical order so output is
	// structured/repeatable
	keys := make([]string, 0)
	for k, _ := range g.Schema.Types {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		t := g.Schema.Types[k]

		switch t.Kind {
		case ast.Scalar:
			g.generateScalarDefinition(scalars, t)
		case ast.Object:
			g.generateTypeDefinition(objects, t)
		case ast.Interface:
			continue
		case ast.Union:
			continue
		case ast.Enum:
			g.generateEnumDefinition(enums, t)
		case ast.InputObject:
			g.generateTypeDefinition(input, t)
		}
	}
}
func (g *Generator) generateScalarDefinition(f *File, t *ast.Definition) error {
	id := getLowerId(t.Name)

	if id != "string" {
		f.Type().Id(id).String().Comment("all scalars are treated as strings")
	}
	f.Func().Id(getScalarContructorName(id)).Params(
		Id("val").String(),
	).Params(
		Id(id),
	).Block(
		Return(Id(id).Parens(Id("val"))),
	)
	return nil
}

func (g *Generator) generateTypeDefinition(f *File, t *ast.Definition) error {
	fields := make([]Code, len(t.Fields))
	for i, v := range t.Fields {
		id := strcase.ToCamel(v.Name)
		tags := map[string]string{
			"graphql": v.Name,
			"json":    fmt.Sprintf("%s,%s", v.Name, "omitempty"),
		}
		comment := fmt.Sprintf("%s %s", v.Name, v.Type.String())
		fields[i] = Id(id).Add(g.getGoType(v.Type)...).Tag(tags).Comment(comment)
	}

	id := getTypeName(t)
	//tags := map[string]string{"graphql": t.Name}
	f.Type().Id(id).Struct(fields...)

	return nil
}

func (g *Generator) generateEnumDefinition(f *File, t *ast.Definition) error {
	enumTypeName := strcase.ToCamel(t.Name)
	f.Type().Id(enumTypeName).Int8()

	values := make([]Code, len(t.EnumValues))

	for i, v := range t.EnumValues {
		// prefix identifier w/ enum type name to prevent collisions
		id := enumTypeName + "_" + strcase.ToCamel(v.Name)
		values[i] = Id(id)

		if i == 0 {
			values[i] = Id(id).Id(enumTypeName).Op("=").Iota()
		}
	}

	f.Const().Defs(values...)

	return nil
}

func getTypeName(t *ast.Definition) string {
	switch t.Kind {
	case ast.Scalar:
		return strcase.ToSnake(t.Name)
	case ast.Object:
		return strcase.ToCamel(t.Name)
	case ast.Interface:
		return strcase.ToCamel(t.Name)
	case ast.Union:
		return strcase.ToCamel(t.Name)
	case ast.Enum:
		return strcase.ToCamel(t.Name)
	case ast.InputObject:
		return strcase.ToSnake(t.Name)
	default:
		return "ERROR"
	}
}

// typeMap is a map of GraphQL scalar types to Go types
var typeMap map[string]string = map[string]string{
	"Boolean": "bool",
	"string":  "string",
	"numeric": "float32",
	"float":   "float32",
	"Int":     "int",
	"uuid":    "string",
	"json":    "map[string]interface{}",
}

func (g *Generator) getGoType(t *ast.Type) (codes []Code) {
	if !t.NonNull {
		codes = append(codes, Op("*"))
	}

	var name string

	if typeIsArray(t) {
		codes = append(codes, Index())
	} else if n, ok := typeMap[t.NamedType]; ok { // build-int go type?
		name = n
	} else {
		// look up the typename in schema and use the same generated type
		// identifier as when creating the type defintion
		for _, td := range g.Schema.Types {
			if td.Name == t.NamedType {
				name = getTypeName(td)
				break
			}
		}
	}

	codes = append(codes, Id(name))
	if t.Elem != nil {
		codes = append(codes, g.getGoType(t.Elem)...)
	}
	return codes
}

func typeIsArray(t *ast.Type) bool {
	return t.NamedType == ""
}

// creates a Go-safe identifier from a source string. Strips keywords
func getLowerId(src string) string {
	id := strcase.ToLowerCamel(src)
	if id == "type" || id == "json" {
		return "_" + id
	}
	return id
}

func getScalarContructorName(name string) string {
	return "New" + strcase.ToCamel(name)
}
