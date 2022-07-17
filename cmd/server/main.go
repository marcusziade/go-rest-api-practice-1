package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/marcusziade/go-rest-api-practice-1/comment"
	"github.com/marcusziade/go-rest-api-practice-1/database"
	"github.com/marcusziade/go-rest-api-practice-1/transport"
)

// App - the struct which contains things like pointers
// to database connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our APP")

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectionParameters := database.ConnectionParameters{
		Username: dbUsername,
		Password: dbPassword,
		Host:     dbHost,
		Database: dbName,
		Port:     dbPort,
	}

	var error error
	newDatabase, error := database.NewDatabase(connectionParameters)
	if error != nil {
		return error
	}

	error = database.MigrateDatabase(newDatabase)
	if error != nil {
		return error
	}

	commentService := comment.NewService(newDatabase)

	handler := transport.NewHandler(commentService)
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
		fmt.Println("Error starting up the REST API")
		fmt.Println(error)
	}
}
