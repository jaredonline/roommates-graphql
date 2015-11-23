package entities

import (
	"log"
	"time"

	// external
	"gopkg.in/gorp.v1"
)

type Expense struct {
	Id            int    `json:"id" db:"id"`
	HouseId       int    `json:"house_id" db:"house_id"`
	AmountInCents int    `json:"amount_in_cents" db:"amount_in_cents"`
	Name          string `json:"name" db:"name"`
	Description   string `json:"description" db:"description"`
	Notes         string `json:"notes" db:"notes"`

	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`

	Loaner Person `json:"loaner"`
	Debts  []Debt `json:"debts"`
}

func GetExpense(dbMap *gorp.DbMap, queryID int) interface{} {
	expense := Expense{}
	err := dbMap.SelectOne(&expense, "SELECT * FROM expenses WHERE id=$1", queryID)
	if err != nil {
		if gorp.NonFatalError(err) {
			log.Print("Ran into some trouble getting person with ID: '", queryID, "': ", err)
		} else {
			log.Fatal("Unable to select user with id '", queryID, "': ", err)
		}
	}
	return expense
}
