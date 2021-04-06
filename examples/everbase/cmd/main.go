package main

import (
	"encoding/json"
	"fmt"

	graphql "github.com/hasura/go-graphql-client"
	"github.com/squk/go-graphql-client-gen/examples/everbase/types"
)

func main() {
	queries := types.Queries{}
	mutations := types.Mutations{}

	c := graphql.NewClient("https://api.production.everbase.co/graphql", nil)
	queries.SetClient(c)
	mutations.SetClient(c)

	fmt.Println("VisitorGeneralInfo Query")
	client, err := queries.VisitorGeneralInfo() // North America
	if err != nil {
		fmt.Println(err)
	}

	b, _ := json.Marshal(client)
	fmt.Println(string(b))

	fmt.Println("ConvertUSDtoEUR Query")
	amount, err := queries.ConvertUSDtoEUR(100.0) // North America
	if err != nil {
		fmt.Println(err)
	}

	b, _ = json.Marshal(amount)
	fmt.Println(string(b))

	fmt.Println("ExchangeRateForVisitorsCurrency Query")
	exc, err := queries.ExchangeRateForVisitorsCurrency() // North America
	if err != nil {
		fmt.Println(err)
	}

	b, _ = json.Marshal(exc)
	fmt.Println(string(b))
}
