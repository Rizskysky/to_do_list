package handlers

import (
	"net/http"
	"to_do_list/module/todo/models"
	"to_do_list/module/todo/repositories"

	"github.com/gin-gonic/gin"
)

// GetTodos :nodoc:
func GetTodos(c *gin.Context) {
	todo, err := repositories.GetAllTodo()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// GetATodo :nodoc:
func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	todo, err := repositories.GetATodo(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// CreateTodo :nodoc:
func CreateTodo(c *gin.Context) {
	var todo models.ToDo
	err := c.BindJSON(&todo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	todo, err = repositories.CreateATodo(todo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// EditTodo :nodoc:
func EditTodo(c *gin.Context) {
	var todo models.ToDo
	id := c.Params.ByName("id")
	todo, err := repositories.GetATodo(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.BindJSON(&todo)
	todo, err = repositories.UpdateTodo(todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// DeleteTodo :nodoc:
func DeleteTodo(c *gin.Context) {
	var todo models.ToDo
	id := c.Params.ByName("id")
	todo, err := repositories.DeleteTodo(todo, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
