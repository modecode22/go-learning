package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("This runs concurrently!")
	}()

	time.Sleep(time.Second)
}
