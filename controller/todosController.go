package main

import (
	"net/http"

	ds "foo.com/dataAccess"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ds.FindAllTodos())
}

func postTodos(c *gin.Context) {
	ds.AddTodo(c.PostForm("text"))
	c.IndentedJSON(http.StatusOK, nil)
}

func delTodos(c *gin.Context) {
	ds.DelTodo(c.PostForm("id"))
	c.IndentedJSON(http.StatusOK, nil)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", postTodos)
	router.DELETE("/todos", delTodos)
	router.Run("localhost:8080")
}
