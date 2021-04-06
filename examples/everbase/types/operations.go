package types

import (
	"context"
	"encoding/json"
	"fmt"

	gographqlclient "github.com/hasura/go-graphql-client"
)

type Queries struct {
	client *gographqlclient.Client
}

func (q *Queries) SetClient(client *gographqlclient.Client) {
	q.client = client
}

type Mutations struct {
	client *gographqlclient.Client
}

func (q *Mutations) SetClient(client *gographqlclient.Client) {
	q.client = client
}
func (q *Queries) VisitorGeneralInfo() (Client, error) {
	query := struct {
		Client struct {
			IpAddress struct {
				Address string `graphql:"address" json:"address,omitempty"`
				City    struct {
					Name       string `graphql:"name" json:"name,omitempty"`
					Population int    `graphql:"population" json:"population,omitempty"`
				} `graphql:"city" json:"city,omitempty"`
				Country struct {
					Name       string `graphql:"name" json:"name,omitempty"`
					Population int    `graphql:"population" json:"population,omitempty"`
					Capital    struct {
						Name string `graphql:"name" json:"name,omitempty"`
					} `graphql:"capital" json:"capital,omitempty"`
					Currencies []struct {
						Name string `graphql:"name" json:"name,omitempty"`
					} `graphql:"currencies" json:"currencies,omitempty"`
				} `graphql:"country" json:"country,omitempty"`
			} `graphql:"ipAddress" json:"ip_address,omitempty"`
		} `graphql:"client" json:"client,omitempty"`
	}{}
	variables := map[string]interface{}{}
	var data Client
	err := q.client.Query(context.Background(), &query, variables)
	if err != nil {
		return data, err
	}
	bytes, err := json.Marshal(query.Client)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func (q *Queries) ConvertUSDtoEUR(amount float32) ([]Currency, error) {
	query := struct {
		Currencies []struct {
			Name  string  `graphql:"name" json:"name,omitempty"`
			ToEUR float32 `graphql:"convert(amount: $amount,to: \"EUR\")" json:"to_eur,omitempty"`
		} `graphql:"currencies(where: {isoCode:{eq:\"USD\"}})" json:"currencies,omitempty"`
	}{}
	variables := map[string]interface{}{"amount": gographqlclient.Float(amount)}
	var data []Currency
	err := q.client.Query(context.Background(), &query, variables)
	if err != nil {
		return data, err
	}
	bytes, err := json.Marshal(query.Currencies)
	if err != nil {
		return data, err
	}
	fmt.Println(string(bytes))
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func (q *Queries) ExchangeRateForVisitorsCurrency() (Client, error) {
	query := struct {
		Client struct {
			IpAddress struct {
				Country struct {
					Name       string `graphql:"name" json:"name,omitempty"`
					Currencies []struct {
						Name  string  `graphql:"name" json:"name,omitempty"`
						ToUSD float32 `graphql:"convert(amount: 1,to: \"USD\")" json:"to_usd,omitempty"`
						ToEUR float32 `graphql:"convert(amount: 1,to: \"EUR\")" json:"to_eur,omitempty"`
					} `graphql:"currencies" json:"currencies,omitempty"`
				} `graphql:"country" json:"country,omitempty"`
			} `graphql:"ipAddress" json:"ip_address,omitempty"`
		} `graphql:"client" json:"client,omitempty"`
	}{}
	variables := map[string]interface{}{}
	var data Client
	err := q.client.Query(context.Background(), &query, variables)
	if err != nil {
		return data, err
	}
	bytes, err := json.Marshal(query.Client)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
