package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Config struct holds application configuration
type Config struct {
	Server ServerConfig
	DB     DATASOURCE
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
	Host string
}

// DBConfig holds database configuration
type DATASOURCE struct {
	Pg   *sql.DB
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbConfig := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUsername, dbPassword, dbName)
	db, err := sql.Open("postgres", dbConfig)
    if err != nil {
        log.Fatal(err)
    }


	cfg := &Config{
		Server: ServerConfig{
			Host: serverHost,
			Port: serverPort,
		},
		DB: DATASOURCE{
			Pg: db,
		},
	}

	return cfg, nil
}
