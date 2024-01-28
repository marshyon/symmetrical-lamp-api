package main

import (
	"fmt"

	"github.com/marshyon/symmetrical-lamp-api/cmd/server/internal/comment"
	"github.com/marshyon/symmetrical-lamp-api/cmd/server/internal/db"
	transportHTTP "github.com/marshyon/symmetrical-lamp-api/cmd/server/internal/transport/http"
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

	// this uses a named import called httpHandler that
	// imports internal/transport/http !!!
	// it is also depencency injected with the comment service
	httphandler := transportHTTP.NewHandler(cmtService)
	if err := httphandler.Serve(); err != nil {
		fmt.Println("Error starting server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println("Error starting up the server")
		fmt.Println(err)
	}

}
