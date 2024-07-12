package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	Id        string `json:"Id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{Id: "1", Item: "clean room", Completed: false},
	{Id: "2", Item: "go to gym", Completed: true},
	{Id: "3", Item: "study go lang", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
}
func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById((id))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}
	context.IndentedJSON(http.StatusOK, todo)
}

func updateTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById((id))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}
func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todo/:id", getTodo)
	router.PATCH("/todo/:id", updateTodoStatus)
	router.POST("/todos", addTodo)
	router.Run("localhost:9090")
}
