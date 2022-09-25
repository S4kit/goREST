package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"Item"`
	Completed bool   `json:"Completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: true},
	{ID: "2", Item: "Make Breakfast", Completed: true},
	{ID: "3", Item: "Make Dinner", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}
func postTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, todos)
}
func getById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todos not found")
}
func ShowToDo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Todo Not Found!"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}
func ToggleToDo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Todo Not Found!"})
		return
	}
	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", ShowToDo)
	router.POST("/todos/:id", ToggleToDo)
	router.POST("/todos", postTodo)
	router.Run("localhost:8080")
}
