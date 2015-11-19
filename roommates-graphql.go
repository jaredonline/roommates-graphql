package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	// external
	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.FieldConfigMap{
			"first_name": &graphql.FieldConfig{
				Type: graphql.String,
			},
			"last_name": &graphql.FieldConfig{
				Type: graphql.String,
			},
			"age": &graphql.FieldConfig{
				Type: graphql.String,
			},
			"id": &graphql.FieldConfig{
				Type: graphql.String,
			},
		},
	},
)

var (
	db  *sql.DB
	err error
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       string `json:"age"`
}

func getUsers() []interface{} {
	users := make([]interface{}, 0)

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			id        int
			firstName string
			lastName  string
			age       int
		)

		rows.Scan(&id, &firstName, &lastName, &age)

		users = append(users, User{
			ID:        fmt.Sprintf("%d", id),
			FirstName: firstName,
			LastName:  lastName,
			Age:       fmt.Sprintf("%d", age),
		})
	}

	return users
}

func main() {
	db, err = sql.Open("postgres", "user=jared dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	fields := graphql.FieldConfigMap{
		"hello": &graphql.FieldConfig{
			Type: graphql.String,
			Resolve: func(p graphql.GQLFRParams) interface{} {
				return "world"
			},
		},
		"foo": &graphql.FieldConfig{
			Type: graphql.String,
			Resolve: func(p graphql.GQLFRParams) interface{} {
				return "bar"
			},
		},
		"users": &graphql.FieldConfig{
			Type: graphql.NewList(userType),
			Resolve: func(p graphql.GQLFRParams) interface{} {
				return getUsers()
			},
		},
		"user": &graphql.FieldConfig{
			Type: userType,
			Resolve: func(p graphql.GQLFRParams) interface{} {
				return getUsers()[0]
			},
		},
	}

	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: fields,
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	// Query
	query := `
		query GetUsers {
			users {
				first_name
				last_name
				age
			}
		}
	`

	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
	}

	r := graphql.Graphql(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute GraphQL operaion: %+v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)

	fmt.Println("%s", string(rJSON))
}
