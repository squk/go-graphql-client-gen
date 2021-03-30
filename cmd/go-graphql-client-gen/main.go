package main

import (
	"fmt"
	"log"
	"os"

	"github.com/squk/go-graphql-client-gen/gen"
	"github.com/urfave/cli/v2"
	//"github.com/squk/go-graphql-client-gen/gen"
)

var schemaFile, schemaUrl, operationsFile, packageName string

func main() {

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "schema",
				Usage:       "load GraphQL schema from `FILE`",
				Hidden:      false,
				Value:       "schema.gql",
				Destination: &schemaFile,
			},
			&cli.StringFlag{
				Name:        "url",
				Usage:       "load GraphQL schema from `URL`",
				Destination: &schemaUrl,
			},
			&cli.StringFlag{
				Name:        "operations",
				Value:       "operations.gql",
				Usage:       "`FILE` of GQL operations to generate code for",
				Destination: &operationsFile,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "package",
				Value:       "types",
				Usage:       "Golang package name to use for generation",
				Destination: &packageName,
				Required:    true,
			},
		},

		Name:  "go-graphql-client-gen",
		Usage: "Generate Golang GraphQL client code",
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "generate files",
			Action: func(c *cli.Context) error {
				if schemaUrl != "" {
					fmt.Println("Schema URL fetching not implemented")
					return nil
				}

				g := gen.NewGenerator(
					gen.WithSchemaFile(schemaFile),
					gen.WithQueriesFile(operationsFile),
				)
				g.Run()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
