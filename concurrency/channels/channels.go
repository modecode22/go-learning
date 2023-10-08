package main

import (
	"fmt"
)

func main() {
	messageChannel := make(chan string)

	go func() {
		messageChannel <- "Hello from a goroutine!"
	}()

	message := <-messageChannel
	fmt.Println(message)
}
