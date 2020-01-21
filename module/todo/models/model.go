package models

type ToDo struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
}

func (b *ToDo) TableName() string {
	return "todo"
}
