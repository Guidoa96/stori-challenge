package model

import "time"

type Transaction struct {
	ID        int
	Amount    float64
	AccountID int
	CreatedAt time.Time
}
