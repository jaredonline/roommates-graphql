package data

import (
	"log"
	// external
	"gopkg.in/gorp.v1"
)

type User struct {
	Id        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Age       int    `json:"age" db:"age"`
}

func GetUserById(dbMap *gorp.DbMap, queryID int) interface{} {
	user := User{}
	err := dbMap.SelectOne(&user, "SELECT * FROM users WHERE id=$1", queryID)
	if err != nil {
		log.Fatal("Unable to select user with id '", queryID, "': ", err)
	}
	return user
}

func GetAllUsers(dbMap *gorp.DbMap) []interface{} {
	var u []User
	_, err := dbMap.Select(&u, "SELECT * FROM users")
	if err != nil {
		log.Fatal("Unable to select all users: ", err)
	}
	users := make([]interface{}, len(u))
	for key, value := range u {
		users[key] = value
	}

	return users
}
