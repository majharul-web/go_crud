package database

import (
	"fmt"
	"go_crud/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "host=aws-1-ap-southeast-1.pooler.supabase.com user=postgres.jveiwkywbmuoflaacqre password=go-crud dbname=postgres port=6543 sslmode=require"
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // Auto-migrate models
    db.AutoMigrate(&models.User{})

    DB = db
    fmt.Println("Database connected successfully")
}
