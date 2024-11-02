package config

import (
	"fmt"
	"log"
	"msvc-syahril-app/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(config *Config) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
        config.DBHost,
        config.DBUser,
        config.DBPassword,
        config.DBName,
        config.DBPort,
        config.DBSSLMode,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    db.AutoMigrate(&models.Project{})

    DB = db
}
