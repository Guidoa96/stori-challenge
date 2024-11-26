package transactions

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"stori/internal/processing/transactions/dao"
)

type TransactionRequest struct {
	Amount  float64 `json:"amount"`
	Account string  `json:"account"`
}

// Handler routes requests based on HTTP method.
func Handler(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	switch r.Method {
	case http.MethodGet:
		GetTransactions(w, r, dbConn)
	case http.MethodPost:
		SaveTransaction(w, r, dbConn)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func SaveTransaction(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var transactionReq TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&transactionReq)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Validate the required fields
	if transactionReq.Amount == 0 || transactionReq.Account == "" {
		http.Error(w, "Missing required fields: amount and account", http.StatusBadRequest)
		return
	}

	err = dao.SaveTransaction(dbConn, transactionReq.Amount, transactionReq.Account)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to save transactionReq: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Transaction saved successfully"))
	w.WriteHeader(http.StatusOK)
}

func GetTransactions(w http.ResponseWriter, r *http.Request, dbConn *sql.DB) {
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

	transactions, err := dao.GetTransactionsForAccount(dbConn, account)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed fetching the account's transaction: %v", err), http.StatusInternalServerError)
		return
	}

	// Define the response structure
	response := map[string]interface{}{
		"status":  "success",
		"message": "Transactions retrieved successfully",
		"data":    transactions,
	}

	// Encode the structured response with indentation for readability
	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}

	// Write the formatted JSON to the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
