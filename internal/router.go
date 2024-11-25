package router

import (
	"net/http"
	"stori/internal/email"
)

func SetupRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// Define routes
	router.HandleFunc("/send_email", email.SendEmailHandler)
	//router.HandleFunc("/")

	return router
}
