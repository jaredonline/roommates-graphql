package data

import (
	"github.com/jaredonline/roommates-graphql/entities"

	// external
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var (
	Person *graphql.Object
)

func init() {
	Person = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Person",
			Fields: graphql.Fields{
				"id":       relay.GlobalIDField("Person", nil),
				"house_id": relay.GlobalIDField("House", nil),
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"updated_at": &graphql.Field{
					Type: graphql.String,
				},
				"balances": &graphql.Field{
					Type: graphql.NewList(Balance),
					Resolve: func(p graphql.ResolveParams) interface{} {
						return entities.GetRoommatesBalances(DbMap, p.Source.(entities.Person))
					},
				},
			},
		},
	)

	Person.AddFieldConfig("roommates", &graphql.Field{
		Type: graphql.NewList(Person),
		Resolve: func(p graphql.ResolveParams) interface{} {
			return entities.GetPersonsRoommates(DbMap, p.Source.(entities.Person))
		},
	})
}
