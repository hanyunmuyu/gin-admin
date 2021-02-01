package seeders

import (
	"fmt"
	"gin-admin/app/models"
	"gin-admin/db"
)

type ActivitySeeder struct {
}

func NewActivitySeeder() *ActivitySeeder {
	return &ActivitySeeder{}
}

func (a ActivitySeeder) Run() {
	for i := 1; i < 20; i++ {
		activity := models.Activity{
			Title:       fmt.Sprintf("title %v", i),
			Description: fmt.Sprintf("description %v", i),
		}
		db.DB.Create(&activity)
	}
}

func (a ActivitySeeder) Drop() {
	db.DB.Exec("DROP TABLE IF EXISTS activities")
}
