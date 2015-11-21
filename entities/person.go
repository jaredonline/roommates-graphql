package entities

import (
	"log"
	"time"

	// external
	"gopkg.in/gorp.v1"
)

type Person struct {
	Id        int       `json:"id" db:"id"`
	HouseId   int       `json:"house_id" db:"house_id"`
	Name      string    `json:"name" db:"name"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Email     string    `json:"email" db:"email"`
	Roommates []Person  `json:"roommates"`
}

func GetPerson(dbMap *gorp.DbMap, queryID int) interface{} {
	user := Person{}
	err := dbMap.SelectOne(&user, "SELECT * FROM people WHERE id=$1", queryID)
	if err != nil {
		if gorp.NonFatalError(err) {
			log.Print("Ran into some trouble getting person with ID: '", queryID, "': ", err)
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
			log.Print("Ran into some trouble getting all people: ", err)
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

func GetPersonsRoommates(dbMap *gorp.DbMap, person Person) []interface{} {
	var u []Person
	_, err := dbMap.Select(&u, "SELECT * FROM people WHERE people.house_id = $1 AND people.id <> $2", person.HouseId, person.Id)
	if err != nil {
		if gorp.NonFatalError(err) {
			log.Print("Ran into some trouble getting all people: ", err)
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
