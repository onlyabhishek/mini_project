
package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=postgres user=postgres password=example dbname=postgres port=5432 sslmode=disable"

	var database *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Connected to PostgreSQL!")
			DB = database
			return
		}
		fmt.Printf("Database not ready. Retrying in 3 seconds... (%d/10)\n", i+1)
		time.Sleep(3 * time.Second)
	}

	log.Fatalf("Failed to connect to database after multiple attempts: %v", err)
}
