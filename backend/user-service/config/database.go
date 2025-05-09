package config

import (
	"fmt"
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil
	}

	// Fetch environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	// Ensure the variables are not empty (for debugging purposes)
	if host == "" || user == "" || password == "" || dbname == "" || port == "" || sslmode == "" {
		return nil, fmt.Errorf("One or more environment variables are missing. Please check your .env file")
	}

	// Construct the DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host, user, password, dbname, port, sslmode)

	// Print the DSN (optional, for debugging)
	fmt.Println("Connecting to database with DSN:", dsn)

	// Connect to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// Print the error for debugging
		fmt.Println("Error connecting to the database:", err)
		return nil, err
	}

	// Enable GORM debugging for SQL logs (optional)
	DB = DB.Debug()

	// Return the database connection
	return DB, nil
}
