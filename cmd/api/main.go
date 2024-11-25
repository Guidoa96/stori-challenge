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
		panic(err)
	}
	defer dbConn.Close()

	// TODO: deleteSave the transaction
	/*err = dao.SaveTransaction(dbConn, transactions[0])
	if err != nil {
		log.Fatalf("Failed to save transaction: %v", err)
	}*/

	// Set up routes
	routes := router.SetupRoutes()

	// Start the server
	fmt.Println("Server starting on :8080")
	err = http.ListenAndServe(":8080", routes)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
