package expenses

import "time"

type Expense struct {
	Date     time.Time
	Amount   float64
	Currency string
	Category string
}
