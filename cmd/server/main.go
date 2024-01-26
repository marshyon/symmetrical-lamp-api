package main

import (
	"context"
	"fmt"

	"github.com/marshyon/symmetrical-lamp-api/cmd/server/internal/comment"
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

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Error migrating db")
		return err
	}

	cmtService := comment.NewService(db)

	cm, err := cmtService.GetComment(
		context.Background(),
		"602d7576-bc5c-11ee-8e50-00155dd54b61",
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("====>[%+v]\n", cm.ID)

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println("Error starting up the server")
		fmt.Println(err)
	}

}
