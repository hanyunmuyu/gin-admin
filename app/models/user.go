package models

type User struct {
	Model
	Password string `json:"password"`
}
