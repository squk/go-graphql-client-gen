package types

type continent_filter_input struct {
	Code *string_query_operator_input `graphql:"code" json:"code,omitempty"` // code StringQueryOperatorInput
}
type country_filter_input struct {
	Code      *string_query_operator_input `graphql:"code" json:"code,omitempty"`           // code StringQueryOperatorInput
	Currency  *string_query_operator_input `graphql:"currency" json:"currency,omitempty"`   // currency StringQueryOperatorInput
	Continent *string_query_operator_input `graphql:"continent" json:"continent,omitempty"` // continent StringQueryOperatorInput
}
type language_filter_input struct {
	Code *string_query_operator_input `graphql:"code" json:"code,omitempty"` // code StringQueryOperatorInput
}
type string_query_operator_input struct {
	Eq    *string    `graphql:"eq" json:"eq,omitempty"`       // eq String
	Ne    *string    `graphql:"ne" json:"ne,omitempty"`       // ne String
	In    *[]*string `graphql:"in" json:"in,omitempty"`       // in [String]
	Nin   *[]*string `graphql:"nin" json:"nin,omitempty"`     // nin [String]
	Regex *string    `graphql:"regex" json:"regex,omitempty"` // regex String
	Glob  *string    `graphql:"glob" json:"glob,omitempty"`   // glob String
}
