package todos

import (
	"fmt"
	"log"

	env "github.com/Netflix/go-env"
	"github.com/lucsky/cuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Environment struct {
	DB struct {
		User     string `env:"DB_User,default=root"`
		Password string `env:"DB_Password"`
		Host     string `env:"DB_Host,default=127.0.0.1"`
		Port     string `env:"DB_Port,default=3306"`
		Database string `env:"Database,default=foo"`
	}
}

type Todo struct {
	gorm.Model
	Tid  string
	Text string
}

var db *gorm.DB

func getDSN() string {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		log.Fatal(err)
	}
	t := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		environment.DB.User, environment.DB.Password, environment.DB.Host, environment.DB.Port, environment.DB.Database)
	return t
}

func getDB() *gorm.DB {

	dsn := getDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Todo{}) // if the table structure not created
	return db
}

// sinple singleton
func getSingleDB() *gorm.DB {
	if db == nil {
		db = getDB()
	}
	return db
}

func AddTodo(t string) {
	db := getSingleDB()
	db.Create(&Todo{Tid: cuid.New(), Text: t})
}

func findTodo(id string) Todo {
	db := getSingleDB()
	var todo Todo
	db.First(&todo, 1)             // find Todo with integer primary key
	db.First(&todo, "Tid = ?", id) // find Todo with code D42
	return todo
}

func DelTodo(id string) {
	db := getSingleDB()
	var todo = findTodo(id)
	db.Delete(&todo, 1)
}

func FindAllTodos() []Todo {
	db := getSingleDB()
	var todos []Todo
	db.Find(&todos)
	return todos
}
