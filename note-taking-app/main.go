package main

import (
	"github.com/gin-gonic/gin"
	"sync"
	"fmt"
)

var notes = make(map[string]string)
var mu sync.Mutex

func main() {
	r := gin.Default()
	r.GET("/", greetHandler)
	r.GET("/greet", greetHandler)
	r.GET("/note", getNotesHandler)
	r.POST("/note", postNoteHandler)

	r.Run(":8080")
}

func greetHandler(c *gin.Context) {
	name := c.DefaultQuery("name", "World")
	c.String(200, "Hello, %s!", name)
}

func getNotesHandler(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	c.JSON(200, notes)
}

func postNoteHandler(c *gin.Context) {
	note := c.Query("note")

	if note == "" {
		c.JSON(400, gin.H{
			"error": "Note content is required",
		})
		return
	}

	key := fmt.Sprintf("%d", len(notes)+1)

	mu.Lock()
	notes[key] = note
	mu.Unlock()

	c.JSON(200, gin.H{
		"message": "Note added successfully!",
	})
}
