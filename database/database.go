package database

import (
	"JCI-Go-API/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	connections, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = connections.AutoMigrate(&models.Topic{}, &models.Comments{})
	if err != nil {
		return
	}
}
