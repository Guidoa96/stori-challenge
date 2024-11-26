package email

import (
	"time"
)

type Summary struct {
	TotalBalance   float64
	Transactions   map[time.Month]int
	AverageCredits map[time.Month]float64
	AverageDebits  map[time.Month]float64
}

// GenerateSummary generates a Summary from a list of transactions.
func GenerateSummary(transactions []TransactionCSV) Summary {
	summary := Summary{
		Transactions:   make(map[time.Month]int),
		AverageCredits: make(map[time.Month]float64),
		AverageDebits:  make(map[time.Month]float64),
	}

	monthCredits := make(map[time.Month]float64)
	monthDebits := make(map[time.Month]float64)
	creditCounts := make(map[time.Month]int)
	debitCounts := make(map[time.Month]int)

	for _, t := range transactions {
		summary.TotalBalance += t.Amount
		month := t.CreatedAt.Month()
		summary.Transactions[month]++

		if t.Amount > 0 {
			monthCredits[month] += t.Amount
			creditCounts[month]++
		} else {
			monthDebits[month] += t.Amount
			debitCounts[month]++
		}
	}

	for month := range summary.Transactions {
		if creditCounts[month] > 0 {
			summary.AverageCredits[month] = monthCredits[month] / float64(creditCounts[month])
		}
		if debitCounts[month] > 0 {
			summary.AverageDebits[month] = monthDebits[month] / float64(debitCounts[month])
		}
	}

	return summary
}
