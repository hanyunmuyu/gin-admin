package models

type Admin struct {
	Model
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	RoleId   uint   `json:"roleId" gorm:"default:0"`
	Password string `json:"password"`
}
