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

func beforeCreate(tx *gorm.DB) (err error) {
	t := models.Topic{}
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

	res := models.Response{
		Code:    201,
		Message: "Success Create Topic",
		Data:    topic,
	}

	return c.Status(http.StatusCreated).JSON(res)
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

	res := models.Response{
		Code:    200,
		Message: "Success get Topic",
		Data: models.Details{
			ID:       topic.ID,
			Title:    topic.Title,
			Content:  topic.Content,
			Comments: comment,
		},
	}
	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) UpdateTopic(c *fiber.Ctx) error {
	id := c.Params("id")
	topic := models.Topic{}

	err := h.db.Where("id = ?", id).First(&topic).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	res := models.Response{
		Code:    200,
		Message: "Success Update Topic",
		Data:    topic,
	}
	h.db.Model(&topic).Updates(topic)
	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) DeleteTopic(c *fiber.Ctx) error {
	id := c.Params("id")
	topic := models.Topic{}

	err := h.db.Where("id = ?", id).First(&topic).Error
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}
	h.db.Delete(&topic)
	res := models.Response{
		Code:    200,
		Message: "Success Delete Topic",
	}
	return c.Status(http.StatusOK).JSON(res)
}
