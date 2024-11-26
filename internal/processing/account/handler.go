package account

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"stori/internal/processing/account/dao"
)

// Handler routes requests based on HTTP method.
func Handler(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	switch r.Method {
	case http.MethodGet:
		GetAccountBalance(w, r, dbConn)
	case http.MethodPost:
		CreateAccount(w, r, dbConn)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetAccountBalance(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	account := r.URL.Query().Get("account_number")
	// Validate the required fields
	if account == "" {
		http.Error(w, "Missing required query param: account", http.StatusBadRequest)
		return
	}

	balance, err := dao.GetAccountBalance(dbConn, account)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed fetching the account's balance: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(balance); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
	}
}

func CreateAccount(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	account := r.URL.Query().Get("account_number")
	// Validate the required fields
	if account == "" {
		http.Error(w, "Missing required query param: account", http.StatusBadRequest)
		return
	}

	_, err := dao.CreateAccount(dbConn, account)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed creating a new account: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Account was created successfully"))
	w.WriteHeader(http.StatusOK)
}
