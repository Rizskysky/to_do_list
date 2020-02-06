package models

import "time"

//ToDo the struct model of database
type ToDo struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Tag         string     `json:"tag"`
	UsersID     string     `json:"user_id"`
	CreatedAt   time.Time  `json:"created_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
