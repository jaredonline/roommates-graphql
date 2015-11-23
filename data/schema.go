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
	Person *graphql.Object

	DbMap *gorp.DbMap
	err   error

	Me entities.Person
)

func init() {
	DbMap, err = InitDB("development")

	Me = entities.GetPerson(DbMap, 1).(entities.Person)

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
			},
		},
	)

	Person.AddFieldConfig("roommates", &graphql.Field{
		Type: graphql.NewList(Person),
		Resolve: func(p graphql.ResolveParams) interface{} {
			return entities.GetPersonsRoommates(DbMap, p.Source.(entities.Person))
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"me": &graphql.Field{
				Type: Person,
				Resolve: func(p graphql.ResolveParams) interface{} {
					return Me
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
