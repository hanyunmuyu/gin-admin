package models

type User struct {
	Model
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
