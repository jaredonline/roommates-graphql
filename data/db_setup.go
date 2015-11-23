package data

import (
	"database/sql"
	"fmt"
	"log"

	// internal
	"github.com/jaredonline/roommates-graphql/entities"

	// external
	"github.com/GeertJohan/go.rice"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Environments map[string]DatabaseEnvironment
}

type DatabaseEnvironment struct {
	Database string `yaml:"database"`
	User     string `yaml:"username"`
	Adapter  string `yaml:"go_adapter"`
}

func InitDB(env string) (*gorp.DbMap, error) {
	config, err := getDBConfig()
	if err != nil {
		log.Fatal("Could not get DB config: ", err)
	}

	// open our db connection
	dbConfig := config.Environments[env]
	db, err := sql.Open(dbConfig.Adapter, fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbConfig.User, dbConfig.Database))
	if err != nil {
		log.Fatal("Could not open database connection: ", err)
	}

	// TODO: this directly calls into postgres
	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbMap.AddTableWithName(entities.Person{}, "people").SetKeys(true, "Id")

	return dbMap, nil
}

func getDBConfig() (DatabaseConfig, error) {
	box, err := rice.FindBox("../db")
	if err != nil {
		log.Fatal("Unable to find static resource box db: ", err)
	}

	data, err := box.String("config.yml")
	if err != nil {
		log.Fatal("Unable to read db/config.yml: ", err)
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
				log.Fatal("Could not marshal database env data: ", err)
			}
			err = yaml.Unmarshal([]byte(d), &dbEnv)
			if err != nil {
				log.Fatal("Could not unmarshal database env data: ", err)
			}
			dbConfig.Environments[envStr] = dbEnv
		}
	}

	return dbConfig, nil
}
