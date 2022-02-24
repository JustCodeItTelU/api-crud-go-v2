package routes

import (
	"JCI-Go-API/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Setup(app *fiber.App) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	h := controllers.NewHandler(db)
	app.Post("/api/createTopic", h.CreateTopic)
	app.Get("/api/getTopics", h.GetTopics)
	app.Get("/api/getTopicID/:id", h.GetTopicByID)
	app.Put("/api/updateTopic/:id", h.UpdateTopic)
	app.Delete("/api/deleteTopic/:id", h.DeleteTopic)

}
