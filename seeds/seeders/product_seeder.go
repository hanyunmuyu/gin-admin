package seeders

import (
	"fmt"
	"gin-admin/app/models"
	"gin-admin/db"
)

type ProductSeeder struct {
}

func NewProductSeeder() *ProductSeeder {
	return &ProductSeeder{}
}

func (p ProductSeeder) Run() {
	for i := 0; i < 20; i++ {
		p := models.Product{Name: fmt.Sprintf("product %d", i)}
		db.DB.Create(&p)
	}
}

func (p ProductSeeder) Drop() {
	db.DB.Exec("DROP TABLE IF EXISTS products")

}
