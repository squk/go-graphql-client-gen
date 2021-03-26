package gen

// The following keywords are reserved and may not be used as identifiers.
// https://golang.org/ref/spec#Keywords
var keywords []string = []string{
	"break",
	"default",
	"func",
	"interface",
	"select",
	"case",
	"defer",
	"go",
	"map",
	"struct",
	"chan",
	"else",
	"goto",
	"package",
	"switch",
	"const",
	"fallthrough",
	"if",
	"range",
	"type",
	"continue",
	"for",
	"import",
	"return",
	"var",

	// technically not keywords
	"string",
	"int",
	"bool",
}

func isGoKeyword(id string) bool {
	for _, kw := range keywords {
		if id == kw {
			return true
		}
	}
	return false
}
