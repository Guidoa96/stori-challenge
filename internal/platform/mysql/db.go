package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	retries = 5
	delay   = 5 * time.Second
)

func Connect() (*sql.DB, error) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
	for i := 0; i < retries; i++ {
		db, err := sql.Open("mysql", dsn)
		if err == nil {
			if pingErr := db.Ping(); pingErr == nil {
				return db, nil
			}
		}
		log.Printf("Retrying database connection... attempt %d/%d", i+1, retries)
		time.Sleep(delay)
	}
	return nil, fmt.Errorf("failed to connect to database after %d attempts", retries)
}
