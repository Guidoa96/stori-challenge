package email

import (
	"fmt"
	"log"
	"net/http"
)

func SendEmailHandler(w http.ResponseWriter, r *http.Request) {
	filePath := "transactions.csv"

	transactions, err := ReadTransactions(filePath)
	if err != nil {
		http.Error(w, "Failed reading transactions.csv", http.StatusInternalServerError)
	}

	summary := GenerateSummary(transactions)
	emailBody, err := FormatSummaryEmail(summary)
	if err != nil {
		err = fmt.Errorf("failed to format email body - %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
