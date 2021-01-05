package models

type Admin struct {
	Model
	Password string `json:"password"`
}
