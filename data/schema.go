package data

import (
	//internal
	"github.com/jaredonline/roommates-graphql/entities"

	//external
	"github.com/graphql-go/graphql"
	"gopkg.in/gorp.v1"
)

var (
	Schema graphql.Schema

	DbMap *gorp.DbMap
	err   error
)

func init() {
	DbMap, err = InitDB("development")

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"me": &graphql.Field{
				Type: Person,
				Resolve: func(p graphql.ResolveParams) interface{} {
					return entities.GetPerson(DbMap, 1)
				},
			},
			"people": &graphql.Field{
				Type: graphql.NewList(Person),
				Resolve: func(p graphql.ResolveParams) interface{} {
					return entities.GetAllPeople(DbMap)
				},
			},
			"person": &graphql.Field{
				Type: Person,
				Resolve: func(p graphql.ResolveParams) interface{} {
					if idQuery, isOK := p.Args["id"].(int); isOK {
						return entities.GetPerson(DbMap, idQuery)
					}
					return nil
				},
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
			},
		},
	})

	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
