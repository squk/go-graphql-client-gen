package types

type Continent struct {
	Code      id        `graphql:"code" json:"code,omitempty"`           // code ID!
	Name      string    `graphql:"name" json:"name,omitempty"`           // name String!
	Countries []Country `graphql:"countries" json:"countries,omitempty"` // countries [Country!]!
}
type Country struct {
	Code      id         `graphql:"code" json:"code,omitempty"`           // code ID!
	Name      string     `graphql:"name" json:"name,omitempty"`           // name String!
	Native    string     `graphql:"native" json:"native,omitempty"`       // native String!
	Phone     string     `graphql:"phone" json:"phone,omitempty"`         // phone String!
	Continent Continent  `graphql:"continent" json:"continent,omitempty"` // continent Continent!
	Capital   *string    `graphql:"capital" json:"capital,omitempty"`     // capital String
	Currency  *string    `graphql:"currency" json:"currency,omitempty"`   // currency String
	Languages []Language `graphql:"languages" json:"languages,omitempty"` // languages [Language!]!
	Emoji     string     `graphql:"emoji" json:"emoji,omitempty"`         // emoji String!
	EmojiU    string     `graphql:"emojiU" json:"emojiU,omitempty"`       // emojiU String!
	States    []State    `graphql:"states" json:"states,omitempty"`       // states [State!]!
}
type Language struct {
	Code   id      `graphql:"code" json:"code,omitempty"`     // code ID!
	Name   *string `graphql:"name" json:"name,omitempty"`     // name String
	Native *string `graphql:"native" json:"native,omitempty"` // native String
	Rtl    bool    `graphql:"rtl" json:"rtl,omitempty"`       // rtl Boolean!
}
type Query struct {
	Continents []Continent `graphql:"continents" json:"continents,omitempty"` // continents [Continent!]!
	Continent  *Continent  `graphql:"continent" json:"continent,omitempty"`   // continent Continent
	Countries  []Country   `graphql:"countries" json:"countries,omitempty"`   // countries [Country!]!
	Country    *Country    `graphql:"country" json:"country,omitempty"`       // country Country
	Languages  []Language  `graphql:"languages" json:"languages,omitempty"`   // languages [Language!]!
	Language   *Language   `graphql:"language" json:"language,omitempty"`     // language Language
	Schema     Schema      `graphql:"__schema" json:"__schema,omitempty"`     // __schema __Schema!
	Type       *Type       `graphql:"__type" json:"__type,omitempty"`         // __type __Type
}
type State struct {
	Code    *string `graphql:"code" json:"code,omitempty"`       // code String
	Name    string  `graphql:"name" json:"name,omitempty"`       // name String!
	Country Country `graphql:"country" json:"country,omitempty"` // country Country!
}
type Directive struct {
	Name        string              `graphql:"name" json:"name,omitempty"`               // name String!
	Description *string             `graphql:"description" json:"description,omitempty"` // description String
	Locations   []DirectiveLocation `graphql:"locations" json:"locations,omitempty"`     // locations [__DirectiveLocation!]!
	Args        []InputValue        `graphql:"args" json:"args,omitempty"`               // args [__InputValue!]!
}
type EnumValue struct {
	Name              string  `graphql:"name" json:"name,omitempty"`                           // name String!
	Description       *string `graphql:"description" json:"description,omitempty"`             // description String
	IsDeprecated      bool    `graphql:"isDeprecated" json:"isDeprecated,omitempty"`           // isDeprecated Boolean!
	DeprecationReason *string `graphql:"deprecationReason" json:"deprecationReason,omitempty"` // deprecationReason String
}
type Field struct {
	Name              string       `graphql:"name" json:"name,omitempty"`                           // name String!
	Description       *string      `graphql:"description" json:"description,omitempty"`             // description String
	Args              []InputValue `graphql:"args" json:"args,omitempty"`                           // args [__InputValue!]!
	Type              Type         `graphql:"type" json:"type,omitempty"`                           // type __Type!
	IsDeprecated      bool         `graphql:"isDeprecated" json:"isDeprecated,omitempty"`           // isDeprecated Boolean!
	DeprecationReason *string      `graphql:"deprecationReason" json:"deprecationReason,omitempty"` // deprecationReason String
}
type InputValue struct {
	Name         string  `graphql:"name" json:"name,omitempty"`                 // name String!
	Description  *string `graphql:"description" json:"description,omitempty"`   // description String
	Type         Type    `graphql:"type" json:"type,omitempty"`                 // type __Type!
	DefaultValue *string `graphql:"defaultValue" json:"defaultValue,omitempty"` // defaultValue String
}
type Schema struct {
	Types            []Type      `graphql:"types" json:"types,omitempty"`                       // types [__Type!]!
	QueryType        Type        `graphql:"queryType" json:"queryType,omitempty"`               // queryType __Type!
	MutationType     *Type       `graphql:"mutationType" json:"mutationType,omitempty"`         // mutationType __Type
	SubscriptionType *Type       `graphql:"subscriptionType" json:"subscriptionType,omitempty"` // subscriptionType __Type
	Directives       []Directive `graphql:"directives" json:"directives,omitempty"`             // directives [__Directive!]!
}
type Type struct {
	Kind          TypeKind      `graphql:"kind" json:"kind,omitempty"`                   // kind __TypeKind!
	Name          *string       `graphql:"name" json:"name,omitempty"`                   // name String
	Description   *string       `graphql:"description" json:"description,omitempty"`     // description String
	Fields        *[]Field      `graphql:"fields" json:"fields,omitempty"`               // fields [__Field!]
	Interfaces    *[]Type       `graphql:"interfaces" json:"interfaces,omitempty"`       // interfaces [__Type!]
	PossibleTypes *[]Type       `graphql:"possibleTypes" json:"possibleTypes,omitempty"` // possibleTypes [__Type!]
	EnumValues    *[]EnumValue  `graphql:"enumValues" json:"enumValues,omitempty"`       // enumValues [__EnumValue!]
	InputFields   *[]InputValue `graphql:"inputFields" json:"inputFields,omitempty"`     // inputFields [__InputValue!]
	OfType        *Type         `graphql:"ofType" json:"ofType,omitempty"`               // ofType __Type
}
