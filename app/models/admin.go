package models

type Admin struct {
	Model
	Name     string `json:"name"`
	RoleId   uint   `json:"roleId" gorm:"default:0"`
	Password string `json:"password"`
}
