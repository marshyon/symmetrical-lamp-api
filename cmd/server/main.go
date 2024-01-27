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

	httphandler := transportHTTP.NewHandler(cmtService)
	if err := httphandler.Serve(); err != nil {
		fmt.Println("Error starting server")
		return err
	}

	// cm, err := cmtService.PostComment(
	// 	context.Background(),
	// 	comment.Comment{
	// 		ID:     "ce598448-bd26-11ee-bf2a-00155dd54b62",
	// 		Slug:   "manual testing 123 ...",
	// 		Author: "Marshyon",
	// 		Body:   "a boo a boo a bibbly booby boo ...",
	// 	},
	// )

	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	// fmt.Printf("result of a created comment ====>[%+v]\n", cm)

	// cm, err = cmtService.GetComment(
	// 	context.Background(),
	// 	"602d7576-bc5c-11ee-8e50-00155dd54b61",
	// )
	// if err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }
	// fmt.Printf("====>[%+v]\n", cm.ID)

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println("Error starting up the server")
		fmt.Println(err)
	}

}
