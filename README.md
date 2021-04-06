# go-graphql-client-gen

# Usage

```sh
# Make a local copy of your GraphQL schema
get-graphql-schema https://my-graphql-endpoint.com/v1/graphql > schema.gql

go-graphql-client-gen --operations operations.sql run
```

# Roadmap

- [x] Type Mapping
- [x] Scalar Generation
- [x] Type Generation
- [x] Object Generation
- [x] Query Generation
- [x] Field Aliasing
- [ ] Inline Fragments
- [x] Fragment Spreads
- [x] Mutation Generation
- [x] CLI commands
- [ ] Remote schema support
- [ ] Custom type mapping/overrides
- [ ] Multiple selection sets in the root operation

# Examples

Visit the [examples](examples/) directory for various working examples.

# Inline Examples

## Countries

Schema from https://countries.trevorblades.com/
Full schema is located at [examples/countries/schema.gql](examples/countries/schema.gql)
Generated types and operations at [examples/countries/types](examples/countries/types)

### Input Schema(truncated for brevity):

```graphql
type Continent {
  code: ID!
  name: String!
  countries: [Country!]!
}

type Country {
  code: ID!
  name: String!
  native: String!
  phone: String!
  continent: Continent!
  capital: String
  currency: String
  languages: [Language!]!
  emoji: String!
  emojiU: String!
  states: [State!]!
}

type Language {
  code: ID!
  name: String
  native: String
  rtl: Boolean!
}

type State {
  code: String
  name: String!
  country: Country!
}
```

### Output Golang

```go
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

type State struct {
	Code    *string `graphql:"code" json:"code,omitempty"`       // code String
	Name    string  `graphql:"name" json:"name,omitempty"`       // name String!
	Country Country `graphql:"country" json:"country,omitempty"` // country Country!
}
```
