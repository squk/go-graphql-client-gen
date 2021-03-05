package main

import (
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"

	. "github.com/dave/jennifer/jen"
)

var schema *ast.Schema

func main() {
	b, err := ioutil.ReadFile("schema.gql")
	if err != nil {
		panic(err)
	}

	src := &ast.Source{
		Name:    "test",
		Input:   string(b),
		BuiltIn: false,
	}

	schema, err = gqlparser.LoadSchema(src)
	generateFromSchema()
}

func generateFromSchema() {
	enums := NewFile("types")
	scalars := NewFile("types")
	input := NewFile("types")
	objects := NewFile("types")

	defer enums.Save("types/enums.go")
	defer scalars.Save("types/scalars.go")
	defer input.Save("types/input.go")
	defer objects.Save("types/objects.go")

	// iterate over map in alphabetical order so output is
	// structured/repeatable
	keys := make([]string, 0)
	for k, _ := range schema.Types {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		t := schema.Types[k]

		switch t.Kind {
		case ast.Scalar:
			id := getTypeName(t)
			scalars.Type().Id(id).String().Comment("all scalars are treated as strings")
		case ast.Object:
			generateTypeDefinition(objects, t)
		case ast.Interface:
			continue
		case ast.Union:
			continue
		case ast.Enum:
			generateEnumDefinition(enums, t)
		case ast.InputObject:
			generateTypeDefinition(input, t)
		}
	}
}

func generateTypeDefinition(f *File, t *ast.Definition) error {
	f.Comment(strcase.ToScreamingSnake(string(t.Kind)))

	fields := make([]Code, len(t.Fields))
	for i, v := range t.Fields {
		id := strcase.ToCamel(v.Name)
		tags := map[string]string{"graphql": v.Name}
		fields[i] = Id(id).Id(getFieldType(v)).Tag(tags)
	}

	id := getTypeName(t)
	//tags := map[string]string{"graphql": t.Name}
	f.Type().Id(id).Struct(fields...)

	return nil
}

func generateEnumDefinition(f *File, t *ast.Definition) error {
	f.Comment(strcase.ToScreamingSnake(string(t.Kind)))

	enumTypeName := strcase.ToCamel(t.Name)
	f.Type().Id(enumTypeName).Int()

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
	"boolean": "bool",
	"string":  "string",
	"numeric": "float32",
	"float":   "float32",
	"Int":     "int",
}

func getFieldType(f *ast.FieldDefinition) string {
	name := f.Type.NamedType

	// look up the typename in schema and use the same generated type
	// identifier as when creating the type defintion
	for _, t := range schema.Types {
		if t.Name == f.Type.NamedType {
			name = getTypeName(t)
			break
		}
	}

	if t, ok := typeMap[name]; ok {
		fmt.Println(t)
		name = t
	}

	return name
}
