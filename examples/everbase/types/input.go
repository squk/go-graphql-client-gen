package types

type city_where struct {
	Id          *where_string `graphql:"id" json:"id,omitempty"`                   // id WhereString
	Name        *where_string `graphql:"name" json:"name,omitempty"`               // name WhereString
	CountryName *where_string `graphql:"countryName" json:"countryName,omitempty"` // countryName WhereString
	Population  *where_float  `graphql:"population" json:"population,omitempty"`   // population WhereFloat
}
type continent_where struct {
	Id         *where_string `graphql:"id" json:"id,omitempty"`                 // id WhereString
	Name       *where_string `graphql:"name" json:"name,omitempty"`             // name WhereString
	GeonamesId *where_int    `graphql:"geonamesId" json:"geonamesId,omitempty"` // geonamesId WhereInt
}
type country_where struct {
	Id         *where_string `graphql:"id" json:"id,omitempty"`                 // id WhereString
	Name       *where_string `graphql:"name" json:"name,omitempty"`             // name WhereString
	Alpha2Code *where_string `graphql:"alpha2Code" json:"alpha2Code,omitempty"` // alpha2Code WhereString
	Alpha3Code *where_string `graphql:"alpha3Code" json:"alpha3Code,omitempty"` // alpha3Code WhereString
	Population *where_int    `graphql:"population" json:"population,omitempty"` // population WhereInt
}
type currency_where struct {
	Id      *where_string `graphql:"id" json:"id,omitempty"`           // id WhereString
	Name    *where_string `graphql:"name" json:"name,omitempty"`       // name WhereString
	IsoCode *where_string `graphql:"isoCode" json:"isoCode,omitempty"` // isoCode WhereString
}
type language_where struct {
	Id         *where_string `graphql:"id" json:"id,omitempty"`                 // id WhereString
	Name       *where_string `graphql:"name" json:"name,omitempty"`             // name WhereString
	Alpha2Code *where_string `graphql:"alpha2Code" json:"alpha2Code,omitempty"` // alpha2Code WhereString
}
type time_zone_where struct {
	Id     *where_string `graphql:"id" json:"id,omitempty"`         // id WhereString
	Name   *where_string `graphql:"name" json:"name,omitempty"`     // name WhereString
	Offset *where_float  `graphql:"offset" json:"offset,omitempty"` // offset WhereFloat
}
type where_float struct {
	Eq  *float32   `graphql:"eq" json:"eq,omitempty"`   // eq Float
	Neq *float32   `graphql:"neq" json:"neq,omitempty"` // neq Float
	In  *[]float32 `graphql:"in" json:"in,omitempty"`   // in [Float!]
	Nin *[]float32 `graphql:"nin" json:"nin,omitempty"` // nin [Float!]
	Lt  *float32   `graphql:"lt" json:"lt,omitempty"`   // lt Float
	Gt  *float32   `graphql:"gt" json:"gt,omitempty"`   // gt Float
}
type where_int struct {
	Eq  *int   `graphql:"eq" json:"eq,omitempty"`   // eq Int
	Neq *int   `graphql:"neq" json:"neq,omitempty"` // neq Int
	In  *[]int `graphql:"in" json:"in,omitempty"`   // in [Int!]
	Nin *[]int `graphql:"nin" json:"nin,omitempty"` // nin [Int!]
	Lt  *int   `graphql:"lt" json:"lt,omitempty"`   // lt Int
	Gt  *int   `graphql:"gt" json:"gt,omitempty"`   // gt Int
}
type where_string struct {
	Eq  *string   `graphql:"eq" json:"eq,omitempty"`   // eq String
	Neq *string   `graphql:"neq" json:"neq,omitempty"` // neq String
	In  *[]string `graphql:"in" json:"in,omitempty"`   // in [String!]
	Nin *[]string `graphql:"nin" json:"nin,omitempty"` // nin [String!]
}
