package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

var todos = []todo{
	{ID: "1", Text: "text1"},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)

	router.Run("localhost:8080")
}
