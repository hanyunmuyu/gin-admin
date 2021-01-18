package models

type Message struct {
	Model
	UserId  uint
	Title   string
	Content string
}
