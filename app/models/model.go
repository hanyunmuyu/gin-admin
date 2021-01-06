package models

import (
	"gin-admin/pkg/utils"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt utils.Time     `json:"createdAt"`
	UpdatedAt utils.Time     `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
