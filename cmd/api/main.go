package main

import (
	"stori/internal/email"
	"stori/internal/processing"
)

func main() {
	// Path to the transaction file
	filePath := "transactions.csv"

	// Read and process the transactions from the csv file
	transactions, err := processing.ReadTransactions(filePath)
	if err != nil {
		panic(err)
	}

	// Generate a summary from the transactions
	summary := processing.GenerateSummary(transactions)

	// Format the email content
	emailBody := email.FormatSummaryEmail(summary)

	err = email.SendEmail("guidoarri96@gmail.com", "Transaction Summary", emailBody)
	if err != nil {
		panic(err)
	}

	println("Email sent successfully!")
}
