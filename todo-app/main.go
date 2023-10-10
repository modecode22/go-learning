package main

import (
	"github.com/gin-gonic/gin"
	"sync"
	"strconv"

)

type TodoItem struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var (
	todoList  []*TodoItem
	nextID    = 1
	todoMutex sync.Mutex
)

func main() {
	r := gin.Default()

	r.GET("/todos", GetTodoList)
	r.POST("/todos", CreateTodo)
	r.PUT("/todos/:id", UpdateTodo)
	r.DELETE("/todos/:id", DeleteTodo)

	r.Run() // Runs on port 8080 by default
}

func GetTodoList(c *gin.Context) {
	todoMutex.Lock()
	defer todoMutex.Unlock()
	c.JSON(200, todoList)
}

func CreateTodo(c *gin.Context) {
	var todo TodoItem
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todoMutex.Lock()
	todo.ID = nextID
	nextID++
	todoList = append(todoList, &todo)
	todoMutex.Unlock()

	c.JSON(201, todo)
}

func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var todo TodoItem
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	todoMutex.Lock()
	for _, t := range todoList {
		if t.ID == id {
			t.Title = todo.Title
			t.Status = todo.Status
			todoMutex.Unlock()
			c.JSON(200, t)
			return
		}
	}
	todoMutex.Unlock()
	c.JSON(404, gin.H{"error": "Todo not found"})
}

func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	todoMutex.Lock()
	for i, t := range todoList {
		if t.ID == id {
			todoList = append(todoList[:i], todoList[i+1:]...)
			todoMutex.Unlock()
			c.JSON(200, gin.H{"message": "Todo deleted"})
			return
		}
	}
	todoMutex.Unlock()
	c.JSON(404, gin.H{"error": "Todo not found"})
}
