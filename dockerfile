FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o transaction-summary cmd/api/main.go

FROM gcr.io/distroless/base-debian11
WORKDIR /app
COPY --from=builder /app/transaction-summary /app/transaction-summary
COPY transactions.csv /app/transactions.csv
COPY logo.png /app/logo.png
COPY .env /app/.env
CMD ["/app/transaction-summary"]
