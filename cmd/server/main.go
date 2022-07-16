package main

import (
	"fmt"
	"net/http"

	"github.com/marcusziade/go-rest-api-practice-1/database"
	"github.com/marcusziade/go-rest-api-practice-1/transport"
)

// App - the struct which contains things like pointers
// to database connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")

	var error error
	database, error := database.NewDatabase()
	if error != nil {
		return error
	}

	println(database.Value)

	handler := transport.NewHandler()
	handler.SetupRoutes()

	if error := http.ListenAndServe(":8080", handler.Router); error != nil {
		fmt.Println("Failed to set up server")
		return error
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if error := app.Run(); error != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(error)
	}
}
