package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"to_do_list/module/todo/models"
	"to_do_list/module/todo/responses"
	"to_do_list/module/todo/requests"
)

//GenerateResponses untuk standar response cakep
func GenerateResponses(c *gin.Context, resp interface{}, err error) {
	response := responses.GenerateResponse(resp, err)
	return
	c.JSON(http.StatusOK, response)
}

//GetTodos handler untuk dapat todos
func GetTodos(c *gin.Context) {
	var todo []models.ToDo
	err := models.GetAllTodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		GenerateResponses(c, todo, err)
	}
}

//GetATodo handler untuk dapetin 1 todo
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

// CreateTodo untuk ngecreate
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

// EditTodo untuk edit todo
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

//DeleteTodo untuk hapus todo
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
