package main

import (
	"encoding/json"
	"fmt"

	graphql "github.com/hasura/go-graphql-client"
	"github.com/squk/go-graphql-client-gen/examples/countries/types"
)

func main() {
	queries := types.Queries{}
	mutations := types.Mutations{}

	c := graphql.NewClient("https://countries.trevorblades.com", nil)
	queries.SetClient(c)
	mutations.SetClient(c)

	fmt.Println("ContinentsFiltered Query")
	continents, err := queries.ContinentsFiltered("152")
	if err != nil {
		fmt.Println(err)
	}
	b, _ := json.Marshal(continents)
	fmt.Println(string(b))
}
