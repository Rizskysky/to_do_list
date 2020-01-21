package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"to_do_list/module/todo/models"
)

func GetTodos(c *gin.Context) {
	var todo []models.ToDo
	err := models.GetAllTodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetATodo(c *gin.Context) {
	var todo models.ToDo
	id := c.Params.ByName("id")
	err := models.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func CreateTodo(c *gin.Context) {
	var todo models.ToDo
	c.BindJSON(&todo)
	err := models.CreateATodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func EditTodo(c *gin.Context) {
	var todo models.ToDo
	id := c.Params.ByName("id")
	err := models.GetATodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.BindJSON(&todo)
	err = models.UpdateTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	var todo models.ToDo
	id := c.Params.ByName("id")
	err := models.DeleteTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
