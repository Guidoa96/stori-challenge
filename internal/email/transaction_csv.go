package email

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

// TransactionCSV is only used to interace with the transactions.csv file.
// In order to interact with the MySql db use model.Transaction instead.
type TransactionCSV struct {
	ID        int
	CreatedAt time.Time
	Amount    float64
}

// ReadTransactions reads the transactions present in a CSV file and parses them into a slice of type TransactionCSV.
func ReadTransactions(filePath string) ([]TransactionCSV, error) {
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

	var transactions []TransactionCSV
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

		transactions = append(transactions, TransactionCSV{
			ID:        id,
			CreatedAt: date,
			Amount:    amount,
		})
	}

	return transactions, nil
}
