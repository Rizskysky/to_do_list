package models

import "time"

//Users the struct users model of database
type Users struct {
	ID        uint       `json:"id"`
	FirstName string     `json:"name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email" binding:"required"`
	Password  string     `json:"password" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
