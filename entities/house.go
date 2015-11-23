package entities

import (
	"log"
	"time"

	// external
	"gopkg.in/gorp.v1"
)

type House struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`

	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	People []Person `json:"people"`
}

func GetHouse(dbMap *gorp.DbMap, queryID int) interface{} {
	house := House{}
	err := dbMap.SelectOne(&house, "SELECT * FROM houses WHERE id=$1", queryID)
	if err != nil {
		if gorp.NonFatalError(err) {
			log.Print("Ran into some trouble getting person with ID: '", queryID, "': ", err)
		} else {
			log.Fatal("Unable to select user with id '", queryID, "': ", err)
		}
	}
	return house
}
