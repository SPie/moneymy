package expenses

import "time"

type Expense struct {
	Date     time.Time `json:"date"`
	Amount   float64   `json:"amount"`
	Category string    `json:"category"`
}

func (e *Expense) Add(exp Expense) {
	e.Amount += exp.Amount
}