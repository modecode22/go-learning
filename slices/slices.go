package main

import "fmt"

func main() {
    fruits := []string{"Apple", "Banana", "Cherry"}

    // Add an item to the slice
    fruits = append(fruits, "Date")

    for _, fruit := range fruits {
        fmt.Println(fruit)
    }
}
