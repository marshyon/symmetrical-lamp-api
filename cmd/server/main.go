package main

import (
	"context"
	"fmt"

	"github.com/marshyon/symmetrical-lamp-api/cmd/server/internal/db"
)

// Run - is going to be responsible for
// instantiating the server and running it
func Run() error {
	fmt.Println("Running server...")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Error connecting to db")
		return err
	}
	if err := db.Ping(context.Background()); err != nil {
		fmt.Println("Error pinging db")
		return err
	}
	fmt.Println("Connected to db")

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println("Error starting up the server")
		fmt.Println(err)
	}

}
