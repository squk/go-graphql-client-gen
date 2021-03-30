package types

import (
	"context"
	"encoding/json"
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
func (q *Queries) ContinentsShallow() ([]Continent, error) {
	query := struct {
		Continents []struct {
			Code id     `graphql:"code" json:"code,omitempty"`
			Name string `graphql:"name" json:"name,omitempty"`
		} `graphql:"continents" json:"continents,omitempty"`
	}{}
	variables := map[string]interface{}{}
	err := q.client.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(query.Continents)
	if err != nil {
		return nil, err
	}
	var data []Continent
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (q *Queries) ContinentsFiltered(continentCode string) ([]Continent, error) {
	query := struct {
		Continents []struct {
			Code      id     `graphql:"code" json:"code,omitempty"`
			Name      string `graphql:"name" json:"name,omitempty"`
			Countries []struct {
				Name string `graphql:"name" json:"name,omitempty"`
			} `graphql:"countries" json:"countries,omitempty"`
		} `graphql:"continents(filter: {code:{eq:$continentCode}})" json:"continents,omitempty"`
	}{}
	variables := map[string]interface{}{"continentCode": gographqlclient.String(continentCode)}
	err := q.client.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(query.Continents)
	if err != nil {
		return nil, err
	}
	var data []Continent
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (q *Queries) ContinentsFilteredDeep(continentCode string) ([]Continent, error) {
	query := struct {
		Continents []struct {
			Code      id     `graphql:"code" json:"code,omitempty"`
			Name      string `graphql:"name" json:"name,omitempty"`
			Countries []struct {
				Code      id      `graphql:"code" json:"code,omitempty"`
				Name      string  `graphql:"name" json:"name,omitempty"`
				Capital   *string `graphql:"capital" json:"capital,omitempty"`
				Languages []struct {
					Code id      `graphql:"code" json:"code,omitempty"`
					Name *string `graphql:"name" json:"name,omitempty"`
				} `graphql:"languages" json:"languages,omitempty"`
				States []struct {
					Code *string `graphql:"code" json:"code,omitempty"`
					Name string  `graphql:"name" json:"name,omitempty"`
				} `graphql:"states" json:"states,omitempty"`
			} `graphql:"countries" json:"countries,omitempty"`
		} `graphql:"continents(filter: {code:{eq:$continentCode}})" json:"continents,omitempty"`
	}{}
	variables := map[string]interface{}{"continentCode": gographqlclient.String(continentCode)}
	err := q.client.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(query.Continents)
	if err != nil {
		return nil, err
	}
	var data []Continent
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (q *Queries) Languages() ([]Language, error) {
	query := struct {
		Languages []struct {
			Code   id      `graphql:"code" json:"code,omitempty"`
			Name   *string `graphql:"name" json:"name,omitempty"`
			Native *string `graphql:"native" json:"native,omitempty"`
		} `graphql:"languages" json:"languages,omitempty"`
	}{}
	variables := map[string]interface{}{}
	err := q.client.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}
	bytes, err := json.Marshal(query.Languages)
	if err != nil {
		return nil, err
	}
	var data []Language
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
