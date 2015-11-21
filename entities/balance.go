package entities

import (
	_ "log"

	// external
	"gopkg.in/gorp.v1"
)

type Balance struct {
	PersonId        int `json:"person_id"`
	ForeignPersonId int `json:"foreign_person_id"`
	Amount          int `json:"amount"`
}

func GetRoommatesBalances(dbMap *gorp.DbMap, person Person) []interface{} {
	//var roommates []Person
	//_, err := dbMap.Select(&roommates, "SELECT id FROM people WHERE people.house_id = $1 AND people.id <> $2", person.HouseId, person.Id)
	//if err != nil {
	//if gorp.NonFatalError(err) {
	//log.Print("Ran into some trouble getting all people: ", err)
	//} else {
	//log.Fatal("Unable to select all users: ", err)
	//}
	//}
	return []interface{}{}
}
