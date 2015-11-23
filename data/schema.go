package data

import (
	//internal
	"github.com/jaredonline/roommates-graphql/entities"

	//external
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
	"gopkg.in/gorp.v1"
)

var (
	Schema graphql.Schema

	Person  *graphql.Object
	Debt    *graphql.Object
	Expense *graphql.Object
	House   *graphql.Object

	DbMap *gorp.DbMap
	err   error

	Me entities.Person
)

func init() {
	DbMap, err = InitDB("development")

	Me = entities.GetPerson(DbMap, 1).(entities.Person)

	House = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "House",
			Fields: graphql.Fields{
				"id": relay.GlobalIDField("House", nil),
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	Person = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Person",
			Fields: graphql.Fields{
				"id": relay.GlobalIDField("Person", nil),
				"name": &graphql.Field{
					Type: graphql.String,
				},
				"email": &graphql.Field{
					Type: graphql.String,
				},
				"updated_at": &graphql.Field{
					Type: graphql.String,
				},
				"house": &graphql.Field{
					Type: House,
					Resolve: func(p graphql.ResolveParams) interface{} {
						return entities.GetHouse(DbMap, p.Source.(entities.Person).HouseId)
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

	//Debt = graphql.NewObject(
	//graphql.ObjectConfig{
	//Name: "Debt",
	//},
	//)

	Expense = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Expense",
			Fields: graphql.Fields{
				"id": relay.GlobalIDField("Expense", nil),
				"name": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"me": &graphql.Field{
				Type: Person,
				Resolve: func(p graphql.ResolveParams) interface{} {
					return Me
				},
			},
		},
	})

	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})
}
