package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"example.com/DB_Go/discord_Table"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

type GormDBWrapper struct {
	*gorm.DB
}

func (db *GormDBWrapper) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	// Get the underlying *sql.DB
	sqlDB, err := db.DB.DB()
	if err != nil {
		return nil, err
	}

	// Call ExecContext on the *sql.DB
	return sqlDB.ExecContext(ctx, query, args...)
}

func (db *GormDBWrapper) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	// Get the underlying *sql.DB
	sqlDB, err := db.DB.DB()
	if err != nil {
		return nil, err
	}

	// Call PrepareContext on the *sql.DB
	return sqlDB.PrepareContext(ctx, query)
}

func (db *GormDBWrapper) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	// Get the underlying *sql.DB
	sqlDB, err := db.DB.DB()
	if err != nil {
		return nil, err
	}

	// Call QueryContext on the *sql.DB
	return sqlDB.QueryContext(ctx, query, args...)
}

func (db *GormDBWrapper) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	// Get the underlying *sql.DB
	sqlDB, err := db.DB.DB()
	if err != nil {
		return nil
	}

	// Call QueryRowContext on the *sql.DB
	return sqlDB.QueryRowContext(ctx, query, args...)
}

func main() {

	//Loading the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Establishing	a connection to the database
	dsn := os.Getenv("DB_URL")
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// implement gormDB.DB() to get the underlying *sql.DB from Gorm wrapper for error-handling & connection checking
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatal(err)
	}

	// ping the database to make sure its connected
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	} else {
		log.Println("Successfully connected to the database.")
	}

	// Initialize API endpoint / router using the Gin web framework
	r := gin.Default()

	// Inititalize query from packet discord_Table
	// Create a GormDBWrapper for gormDB and point it to discord_Table/db.go/New so gormDB can work with DBTX interface
	queries := discord_Table.New(&GormDBWrapper{gormDB})

	r.GET("/api/v1/discord-events", func(c *gin.Context) {
		// Get the 'offset' and 'limit' query parameters
		offsetStr := c.DefaultQuery("offset", "50")
		limitStr := c.DefaultQuery("limit", "0")

		// Convert 'offset' and 'limit' to integers
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset value"})
			return
		}
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
			return
		}

		// Get a subset of Discord Gateway Events from the database
		events, err := queries.GetDiscordGatewayEvents(context.Background(), limit, offset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Return the events as a JSON response
		c.JSON(http.StatusOK, events)
	})

	// Start the server
	r.Run()

}
