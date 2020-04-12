package handlers

import (
	"fmt"
	"net/http"
	"to_do_list/module/todo/models"
	"to_do_list/module/todo/repositories"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	todoRepo repositories.Todo
}

func TodoHandler(todo repositories.Todo) *todoHandler {
	return &todoHandler{
		todoRepo: todo,
	}

}

// GetTodos :nodoc:
func (handler *todoHandler) GetTodos(c *gin.Context) {
	todo, err := handler.todoRepo.GetAllTodo()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// GetATodo :nodoc:
func (handler *todoHandler) GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	todo, err := handler.todoRepo.GetATodo(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// CreateTodo :nodoc:
func (handler *todoHandler) CreateTodo(c *gin.Context) {
	var todo models.ToDo
	requestPayload := make(map[string]interface{})

	err := c.BindJSON(&requestPayload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	fmt.Println(requestPayload)
	todo, err = handler.todoRepo.CreateATodo(todo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// EditTodo :nodoc:
func (handler *todoHandler) EditTodo(c *gin.Context) {
	var todo models.ToDo
	id := c.Params.ByName("id")
	todo, err := handler.todoRepo.GetATodo(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.BindJSON(&todo)
	todo, err = handler.todoRepo.UpdateTodo(todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// DeleteTodo :nodoc:
func (handler *todoHandler) DeleteTodo(c *gin.Context) {
	var todo models.ToDo
	id := c.Params.ByName("id")
	todo, err := handler.todoRepo.DeleteTodo(todo, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
