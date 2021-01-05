package models

import (
	"gin-admin/pkg/utils"
)

type Model struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt utils.Time `json:"createdAt"`
	UpdatedAt utils.Time `json:"updatedAt"`
	DeletedAt utils.Time `gorm:"index" json:"deletedAt"`
}
