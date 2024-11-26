package email

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

/*// FormatSummaryEmail formats the transaction summary into a string suitable for email content.
func FormatSummaryEmail(summary Summary) string {
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
}*/

// FormatSummaryEmail formats the transaction summary into an HTML string suitable for email content.
func FormatSummaryEmail(summary Summary) (string, error) {
	// Read and encode the logo as a Base64 string
	logoPath := "logo.png" // Path to the logo file
	logoBytes, err := os.ReadFile(logoPath)
	if err != nil {
		return "", fmt.Errorf("failed to read logo file: %w", err)
	}
	logoBase64 := base64.StdEncoding.EncodeToString(logoBytes)

	var builder strings.Builder

	builder.WriteString(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Transaction Summary</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            background-color: #f9f9f9;
            color: #333;
            margin: 0;
            padding: 0;
        }
        .container {
            width: 90%;
            max-width: 600px;
            margin: 20px auto;
            background: #fff;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        }
        .header {
            text-align: center;
            margin-bottom: 20px;
        }
        .header img {
            max-width: 150px;
            height: auto;
        }
        .summary {
            margin-bottom: 20px;
        }
        .month {
            background-color: #f4f4f4;
            padding: 10px;
            border-radius: 4px;
            margin-bottom: 10px;
        }
        .month h4 {
            margin: 0;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <img src="data:image/png;base64,` + logoBase64 + `" alt="Stori Logo">
        </div>
        <div class="summary">
            <h2>Transaction Summary</h2>
            <p><strong>Total Balance:</strong> $` + fmt.Sprintf("%.2f", summary.TotalBalance) + `</p>
        </div>
`)

	for month, count := range summary.Transactions {
		builder.WriteString(`
        <div class="month">
            <h4>Month: ` + month.String() + `</h4>
            <p>Number of Transactions: ` + fmt.Sprintf("%d", count) + `</p>
            <p>Average Credit: $` + fmt.Sprintf("%.2f", summary.AverageCredits[month]) + `</p>
            <p>Average Debit: $` + fmt.Sprintf("%.2f", summary.AverageDebits[month]) + `</p>
        </div>
`)
	}

	builder.WriteString(`
    </div>
</body>
</html>
`)

	return builder.String(), nil
}
