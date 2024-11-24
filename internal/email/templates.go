package email

import (
	"fmt"
	"stori/internal/processing"
	"strings"
)

// FormatSummaryEmail formats the transaction summary into a string suitable for email content.
func FormatSummaryEmail(summary processing.Summary) string {
	var builder strings.Builder

	builder.WriteString("Transaction Summary:\n")
	builder.WriteString(fmt.Sprintf("Total Balance: %.2f\n", summary.TotalBalance))

	for month, count := range summary.Transactions {
		builder.WriteString(fmt.Sprintf("Month: %s\n", month))
		builder.WriteString(fmt.Sprintf("Number of Transactions: %d\n", count))
		builder.WriteString(fmt.Sprintf("Average Credit: %.2f\n", summary.AverageCredits[month]))
		builder.WriteString(fmt.Sprintf("Average Debit: %.2f\n\n", summary.AverageDebits[month]))
	}

	return builder.String()
}
