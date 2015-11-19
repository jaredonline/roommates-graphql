package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// external
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
}

func InitDB(env string) (*gorp.DbMap, error) {
	config, err := getDBConfig()

	// open our db connection
	dbConfig := config.Environments[env]
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbConfig.User, dbConfig.Database))
	if err != nil {
		log.Fatal("Could not open database connection: ", err)
	}

	dbMap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbMap.AddTableWithName(Person{}, "people").SetKeys(true, "Id")

	return dbMap, nil
}

func getDBConfig() (DatabaseConfig, error) {
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

	return dbConfig, nil
}
