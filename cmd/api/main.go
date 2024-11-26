package main

import (
	"fmt"
	"log"
	"net/http"
	"stori/internal"
	"stori/internal/platform/mysql"
)

func main() {
	// Connect to the database
	dbConn, err := mysql.Connect()
	if err != nil {
		log.Fatalf("Error starting mysql DB: %s", err)
	}
	defer dbConn.Close()

	// Set up routes
	routes := router.SetupRoutes(dbConn)

	// Start the server
	fmt.Println("Server starting on :8080")
	err = http.ListenAndServe(":8080", routes)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
