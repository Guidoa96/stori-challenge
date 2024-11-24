package processing

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type Transaction struct {
	Id     int
	Date   time.Time
	Amount float64
}

// ReadTransactions reads the transactions present in a CSV file and parses them into a slice of type Transaction.
func ReadTransactions(filePath string) ([]Transaction, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Skip the first row (header)
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var transactions []Transaction
	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		date, err := time.Parse("2006-01-02", record[1]) // Assume date format is YYYY-MM-DD
		if err != nil {
			return nil, err
		}

		amount, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, Transaction{
			Id:     id,
			Date:   date,
			Amount: amount,
		})
	}

	return transactions, nil
}
