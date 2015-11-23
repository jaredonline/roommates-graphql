package entities

import (
	"log"
	"time"

	// external
	"gopkg.in/gorp.v1"
)

type Debt struct {
	Id            int    `json:"id" db:"id"`
	AmountInCents int    `json:"amount_in_cents" db:"amount_in_cents"`
	Person        Person `json:"person" db:"person"`
	Loaner        Person `json:"loaner" db:"loaner"`

	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func GetDebt(dbMap *gorp.DbMap, queryID int) interface{} {
	debt := Debt{}
	err := dbMap.SelectOne(&debt, "SELECT * FROM debts WHERE id=$1", queryID)
	if err != nil {
		if gorp.NonFatalError(err) {
			log.Print("Ran into some trouble getting person with ID: '", queryID, "': ", err)
		} else {
			log.Fatal("Unable to select user with id '", queryID, "': ", err)
		}
	}
	return debt
}
