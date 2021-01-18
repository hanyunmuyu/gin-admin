package seeders

import (
	"fmt"
	"gin-admin/app/models"
	"gin-admin/db"
)

type MessageSeeder struct {
}

func (m MessageSeeder) Run() {
	for i := 1; i < 20; i++ {
		message := models.Message{
			UserId:  uint(i),
			Title:   fmt.Sprintf("title %v", i),
			Content: fmt.Sprintf("message %v", i),
		}
		db.DB.Create(&message)
	}
}

func (m MessageSeeder) Drop() {
	db.DB.Exec("DROP TABLE IF EXISTS messages")
}

func NewMessageSeeder() *MessageSeeder {
	return &MessageSeeder{}
}
