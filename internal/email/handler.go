package email

import (
	"fmt"
	"log"
	"net/http"
	"stori/internal/processing"
)

func SendEmailHandler(w http.ResponseWriter, r *http.Request) {
	filePath := "transactions.csv"

	transactions, err := processing.ReadTransactions(filePath)
	if err != nil {
		http.Error(w, "Failed reading transactions.csv", http.StatusInternalServerError)
	}

	summary := processing.GenerateSummary(transactions)
	emailBody := FormatSummaryEmail(summary)

	err = SendEmail("client@transaction.com", "Transaction Summary", emailBody)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte("Email sent successfully")); err != nil {
		fmt.Printf("Failed to write response: %v\n", err)
	}
}
