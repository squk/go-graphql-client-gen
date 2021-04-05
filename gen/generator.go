package gen

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/iancoleman/strcase"
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

func (g *Generator) Run() {
	path := g.Package
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModeDir)
	}

	g.generateTypes()
	g.generateOperations()
}

func (g *Generator) generateTypes() {
	enums := NewFile(g.Package)
	scalars := NewFile(g.Package)
	input := NewFile(g.Package)
	objects := NewFile(g.Package)

	defer enums.Save(filepath.Join(g.Package, "enums.go"))
	defer scalars.Save(filepath.Join(g.Package, "scalars.go"))
	defer input.Save(filepath.Join(g.Package, "input.go"))
	defer objects.Save(filepath.Join(g.Package, "objects.go"))

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

	if isGoKeyword(id) {
		return nil
	}

	// generate alias to string
	f.Type().Id(id).String().Comment("all scalars are treated as strings")

	// geneate constructor for out of package access
	f.Func().Id(getScalarConstructorName(id)).Params(
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

func typeIsArray(t *ast.Type) bool {
	return t.NamedType == ""
}
