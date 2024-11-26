package dao

import (
	"database/sql"
	"errors"
	"stori/internal/model"
)

/*type TransactionsDao interface {
	// SaveTransaction Creates a new snapshot in the kvs
	SaveTransaction(db *sql.DB, transaction processing.Transaction) error
	// UpdateAccountBalance Gets the last snapshot from the kvs
	UpdateAccountBalance(db *sql.DB, newBalance float64) error
}

type transactionsDao struct {
	db *sql.DB
}*/

// SaveTransaction inserts a transaction into the database for the given account
// and updates the total_balance in the accounts table.
func SaveTransaction(db *sql.DB, amount float64, account string) error {
	// Start a db transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Check if the account exists and retrieve its ID and current balance
	var accountID int
	var currentBalance float64
	err = tx.QueryRow("SELECT id, total_balance FROM accounts WHERE account_number = ?", account).Scan(&accountID, &currentBalance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("account not found")
		}
		return err
	}

	// Insert the transaction
	_, err = tx.Exec("INSERT INTO transactions (amount, account_id) VALUES (?, ?)", amount, accountID)
	if err != nil {
		return err
	}

	// Update the account's total balance
	newBalance := currentBalance + amount
	_, err = tx.Exec("UPDATE accounts SET total_balance = ? WHERE id = ?", newBalance, accountID)
	if err != nil {
		return err
	}

	return nil
}

// GetTransactionsForAccount retrieves all transactions for the specified account.
func GetTransactionsForAccount(db *sql.DB, account string) ([]model.Transaction, error) {
	var transactions []model.Transaction

	// Query to get account ID
	var accountID int
	err := db.QueryRow("SELECT id FROM accounts WHERE account_number = ?", account).Scan(&accountID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("account not found")
		}
		return nil, err
	}

	// Query to retrieve transactions for the account
	rows, err := db.Query("SELECT id, amount, account_id, created_at FROM transactions WHERE account_id = ? ORDER BY created_at DESC", accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Process each transaction
	for rows.Next() {
		var transaction model.Transaction
		err := rows.Scan(&transaction.ID, &transaction.Amount, &transaction.AccountID, &transaction.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
