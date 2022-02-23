package routes

import (
	"JCI-Go-API/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	h := controllers.Handler{}
	app.Post("/api/createTopic", h.CreateTopic)
	app.Get("/api/getTopics", h.GetTopics)
	app.Get("/api/getTopicID/{id}", h.GetTopicByID)
	app.Put("/api/updateTopic/{id}", h.UpdateTopic)
	app.Delete("/api/deleteTopic/{id}", h.DeleteTopic)

}
