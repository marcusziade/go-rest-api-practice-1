package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ConnectionParameters struct {
	Username string
	Password string
	Host     string
	Database string
	Port     string
}

// NewDatabase - returns a pointer to a new database connection
func NewDatabase(parameters ConnectionParameters) (*gorm.DB, error) {

	// dbConnectionString := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + strconv.Itoa(dbPort) + ")/" + dbTable
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", parameters.Host, parameters.Port, parameters.Username, parameters.Database, parameters.Password)
	fmt.Println(connectionString)

	database, error := gorm.Open("postgres", connectionString)
	if error != nil {
		return database, error
	}

	if error := database.DB().Ping(); error != nil {
		return database, error
	}

	return database, nil
}
