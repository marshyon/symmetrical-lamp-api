package main

import "fmt"

// Run - is going to be responsible for
// instantiating the server and running it
func Run() error {
	fmt.Println("Running server...")
	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println("Error starting up the server")
		fmt.Println(err)
	}

}
