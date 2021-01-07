package models

type Permission struct {
	Model
	ApiPath  string `json:"apiPath"`
	Rule     string `json:"rule"`
	Method   string `json:"method"`
	Title    string `json:"title"`
	ParentId uint   `json:"parentId"`
	IsMenu   uint8  `json:"isMenu"`
	Path     string `json:"path"`
}
