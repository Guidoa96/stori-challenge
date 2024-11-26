package dao

import (
	"database/sql"
	"errors"
)

// GetAccountBalance retrieves the total balance of the specified account.
func GetAccountBalance(db *sql.DB, account string) (float64, error) {
	var balance float64

	// Query the total balance for the given account number
	err := db.QueryRow("SELECT total_balance FROM accounts WHERE account_number = ?", account).Scan(&balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("account not found")
		}
		return 0, err
	}

	return balance, nil
}

// CreateAccount creates a new account with the specified account number and initial balance.
func CreateAccount(db *sql.DB, accountNumber string) (int64, error) {
	// Insert a new account into the accounts table
	result, err := db.Exec("INSERT INTO accounts (account_number, total_balance) VALUES (?, ?)", accountNumber, 0)
	if err != nil {
		return 0, err
	}

	// Retrieve the ID of the newly created account
	accountID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return accountID, nil
}
