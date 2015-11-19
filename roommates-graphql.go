package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	// external
	"github.com/graphql-go/graphql"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"first_name": &graphql.Field{
				Type: graphql.String,
			},
			"last_name": &graphql.Field{
				Type: graphql.String,
			},
			"age": &graphql.Field{
				Type: graphql.Int,
			},
			"id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

var (
	db  *sql.DB
	err error
)

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
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
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			Age:       age,
		})
	}

	return users
}

type DatabaseConfig struct {
	Environments map[string]DatabaseEnvironment
}

type DatabaseEnvironment struct {
	Database string `yaml:"database"`
	User     string `yaml:"username"`
}

func main() {
	// get our db config from the yml file
	file, err := os.Open("db/config.yml")
	if err != nil {
		log.Fatal("Could not open db/config.yml: ", err)
	}
	fi, err := file.Stat()
	if err != nil {
		log.Fatal("Could not stat db/config.yml: ", err)
	}

	data := make([]byte, fi.Size())
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	rawConf := make(map[interface{}]interface{})
	dbConfig := DatabaseConfig{
		Environments: make(map[string]DatabaseEnvironment),
	}

	err = yaml.Unmarshal([]byte(data), &rawConf)
	if err != nil {
		log.Fatal(err)
	}

	for env, c := range rawConf {
		if envStr, ok := env.(string); ok {
			dbEnv := DatabaseEnvironment{}
			d, err := yaml.Marshal(&c)
			if err != nil {
			}
			err = yaml.Unmarshal([]byte(d), &dbEnv)
			if err != nil {
			}
			dbConfig.Environments[envStr] = dbEnv
		}
	}

	// open our db connection
	db, err = sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbConfig.Environments["development"].User, dbConfig.Environments["development"].Database))
	if err != nil {
		log.Fatal("Could not open database connection: ", err)
	}

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) interface{} {
				return "world"
			},
		},
		"foo": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) interface{} {
				return "bar"
			},
		},
		"users": &graphql.Field{
			Type: graphql.NewList(userType),
			Resolve: func(p graphql.ResolveParams) interface{} {
				return getUsers()
			},
		},
		"user": &graphql.Field{
			Type: userType,
			Resolve: func(p graphql.ResolveParams) interface{} {
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

	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute GraphQL operaion: %+v", r.Errors)
	}

	rJSON, _ := json.Marshal(r)

	fmt.Println("%s", string(rJSON))
}
