package dao

import (
	"database/sql"
	"stori/internal/processing"
)

// LastSnapshotDao defines the contract with the methods
// that values domain needs to know
type TransactionsDao interface {
	// CreateSnapshot Creates a new snapshot in the kvs
	SaveTransaction(db *sql.DB, transaction processing.Transaction) error
	// GetLastSnapshot Gets the last snapshot from the kvs
	UpdateAccountBalance(db *sql.DB, newBalance float64) error
}

type transactionsDao struct {
	db *sql.DB
}

func SaveTransaction(db *sql.DB, transaction processing.Transaction) error {
	query := "INSERT INTO transactions (transaction_date, amount) VALUES (?, ?)"
	_, err := db.Exec(query, transaction.Date, transaction.Amount)
	return err
}

func UpdateAccountBalance(db *sql.DB, newBalance float64) error {
	query := "INSERT INTO accounts (total_balance) VALUES (?)"
	_, err := db.Exec(query, newBalance)
	return err
}
