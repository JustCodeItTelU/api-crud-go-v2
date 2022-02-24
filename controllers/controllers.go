package controllers

import (
	"JCI-Go-API/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
	}
}

type Topic struct {
	models.Topic
}

func (t *Topic) beforeCreated(tx *gorm.DB) (err error) {
	if t.ID == uuid.Nil {
		uuid := uuid.New()
		t.ID = uuid
	}
	return nil
}

func (h *Handler) CreateTopic(c *fiber.Ctx) error {
	topic := models.Topic{}
	if err := c.BodyParser(&topic); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	err := h.db.Create(&topic).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "success create topic",
		"data":    topic,
	})
}

func (h *Handler) GetTopics(c *fiber.Ctx) error {
	topic := models.Topic{}
	err := h.db.Find(&topic).Error
	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(topic)

}

func (h *Handler) GetTopicByID(c *fiber.Ctx) error {
	id := c.Params("id")
	topic := models.Topic{}
	comment := models.Comments{}

	err := h.db.Where("id = ?", id).First(&topic).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message":  "success get Topic",
		"Data":     topic,
		"Comments": comment,
	})
}

func (h *Handler) UpdateTopic(c *fiber.Ctx) error {
	id := c.Params("id")
	topic := models.Topic{}

	err := h.db.Where("id = ?", id).First(&topic).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	h.db.Model(&topic).Updates(topic)
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Success Update Topic",
		"Data":    topic,
	})
}

func (h *Handler) DeleteTopic(c *fiber.Ctx) error {
	id := c.Params("id")
	topic := models.Topic{}

	err := h.db.Where("id = ?", id).First(&topic).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	h.db.Delete(&topic)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Message": "Success Delete Topic",
	})
}
