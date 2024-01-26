package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase() (*Database, error) {
	ConnectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)
	fmt.Printf("connection string: %s\n", ConnectionString)
	// comments-api  | connection string:
	// host=db port=5432 user= dbname= password=postgres sslmode=disable
	db, err := sqlx.Connect("postgres", ConnectionString)
	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to db : %s", err)
	}

	return &Database{Client: db}, nil

}

func (d *Database) Ping(ctx context.Context) error {
	// return d.Client.DB.Ping(ctx)
	return d.Client.Ping()
}
