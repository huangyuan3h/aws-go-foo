package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucsky/cuid"
)

type todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

var todos = []todo{
	{ID: cuid.New(), Text: "text1"},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func postTodos(c *gin.Context) {

	todos = append(todos, todo{cuid.New(), c.PostForm("text")})
	c.IndentedJSON(http.StatusOK, todos)
}

func delTodos(c *gin.Context) {
	filtered := []todo{}
	deletedID := c.PostForm("id")
	for _, item := range todos {
		if item.ID != deletedID {
			filtered = append(filtered, item)
		}
	}
	todos = filtered
	c.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", postTodos)
	router.DELETE("/todos", delTodos)
	router.Run("localhost:8080")
}
