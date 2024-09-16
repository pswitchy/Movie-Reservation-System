package db

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("Failed to connect to the database:", err)
        panic(err)
    }

    fmt.Println("Successfully connected to the database!")
}
