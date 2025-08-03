package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

func TestPostgresConnection(t *testing.T) {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		t.Fatalf("failed to open connection: %v", err)
	}
	defer db.Close()

	// Ping database to verify connection
	if err := db.Ping(); err != nil {
		t.Fatalf("failed to ping database: %v", err)
	}
}
