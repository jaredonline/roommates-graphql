package data

import (
	"log"
	"time"
	// external
	"gopkg.in/gorp.v1"
)

type Person struct {
	Id        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Email     string    `json:"email" db:"email"`
}

func GetPerson(dbMap *gorp.DbMap, queryID int) interface{} {
	user := Person{}
	err := dbMap.SelectOne(&user, "SELECT * FROM users WHERE id=$1", queryID)
	if err != nil {
		if gorp.NonFatalError(err) {
			log.Print("Unable to select user with id '", queryID, "': ", err)
		} else {
			log.Fatal("Unable to select user with id '", queryID, "': ", err)
		}
	}
	return user
}

func GetAllPeople(dbMap *gorp.DbMap) []interface{} {
	var u []Person
	_, err := dbMap.Select(&u, "SELECT * FROM people")
	if err != nil {
		if gorp.NonFatalError(err) {
			log.Print("Unable to select all users: ", err)
		} else {
			log.Fatal("Unable to select all users: ", err)
		}
	}
	users := make([]interface{}, len(u))
	for key, value := range u {
		users[key] = value
	}

	return users
}
