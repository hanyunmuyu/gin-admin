package models

type Product struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
