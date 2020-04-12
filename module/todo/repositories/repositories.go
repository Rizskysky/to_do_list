package repositories

import (
	"to_do_list/module/todo/models"

	"github.com/jinzhu/gorm"
)

//TodoRepository :nodoc:
type TodoRepository struct {
	db *gorm.DB
}

//Todo interface for all
type Todo interface {
	GetAllTodo() ([]models.ToDo, error)
	CreateATodo(todo models.ToDo) (models.ToDo, error)
	GetATodo(id string) (models.ToDo, error)
	UpdateTodo(todo models.ToDo, id string) (models.ToDo, error)
	DeleteTodo(todo models.ToDo, id string) (models.ToDo, error)
}
