package models

import (
	"github.com/google/uuid"
	"time"
)

type (
	Topic struct {
		ID        uuid.UUID
		Title     string `json:"title"`
		Content   string `json:"content"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	Details struct {
		ID       uuid.UUID
		Title    string      `json:"title"`
		Content  string      `json:"content"`
		Comments interface{} `json:"comments"`
	}
	Comments struct {
		ID      uuid.UUID
		Comment string `json:"comment"`
		IDTopic int    `json:"id_topic"`
	}
	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)
