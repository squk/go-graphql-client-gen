package gen

import (
	"github.com/iancoleman/strcase"
	"github.com/vektah/gqlparser/v2/ast"

	. "github.com/dave/jennifer/jen"
)

// creates a Go-safe identifier from a source string. Strips keywords
func getLowerId(src string) string {
	id := strcase.ToLowerCamel(src)
	if id == "type" || id == "json" { // TODO: make keywords more flexible
		return "_" + id
	}
	return id
}

// typeMap is a map of GraphQL scalar types to Go types
var typeMap map[string]string = map[string]string{
	"Boolean": "bool",
	"string":  "string",
	"numeric": "float32",
	"float":   "float32",
	"Int":     "int",
	"json":    "map[string]interface{}",
}

func isMappableScalar(id string) bool {
	_, ok := typeMap[id]
	return ok
}

// getGoType crawls the AST and returns a Go equivalent type as Jennifer code
func (g *Generator) getGoType(t *ast.Type) (codes []Code) {
	if !t.NonNull {
		codes = append(codes, Op("*"))
	}

	var name string

	if typeIsArray(t) {
		codes = append(codes, Index())
	} else if n, ok := typeMap[t.NamedType]; ok { // build-in go type?
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

// getTypeName converts an ast.Definition to a string based of its kind to
// match Goland style casing
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

// Create name for Scalar constructors. e.g. NewUuid()
func getScalarConstructorName(name string) string {
	return "New" + strcase.ToCamel(name)
}
