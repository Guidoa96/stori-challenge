package router

import (
	"database/sql"
	"net/http"
	"stori/internal/email"
	"stori/internal/processing/account"
	"stori/internal/processing/transactions"
)

func SetupRoutes(dbConn *sql.DB) *http.ServeMux {
	router := http.NewServeMux()

	// Define routes
	router.HandleFunc("/send_email", email.SendEmailHandler)
	router.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		transactions.Handler(w, r, dbConn)
	})
	router.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		account.Handler(w, r, dbConn)
	})

	return router
}
