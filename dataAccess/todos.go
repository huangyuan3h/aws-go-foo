package main

import (
	"github.com/lucsky/cuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	id   string
	text string
}

var db *gorm.DB

func getDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/foo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// sinple singleton
func getSingleDB() *gorm.DB {
	if db == nil {
		db = getDB()
	}
	return db
}

func addTodo(t string) {
	db := getSingleDB()
	db.Create(&Todo{id: cuid.New(), text: t})
}

func findTodo(id string) Todo {
	db := getSingleDB()
	var todo Todo
	db.First(&todo, 1)               // find Todo with integer primary key
	db.First(&todo, "id = ?", "D42") // find Todo with code D42
	return todo
}

func delTodo(id string) {
	db := getSingleDB()
	var todo = findTodo(id)
	db.Delete(&todo, 1)
}

func findAllTodos() []Todo {
	db := getSingleDB()
	var todos []Todo
	db.Find(&todos)
	return todos
}
