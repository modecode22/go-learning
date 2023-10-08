package main
import "fmt"
// this is a struct
type Person struct {
    Name   string
    Age    int
    Height float64
}

// this is a method 
func (p Person) Speak() string {
    return "Hello, my name is " + p.Name
}

type Speaker interface {
    Speak() string
}

func Introduce(speaker Speaker) {
    fmt.Println(speaker.Speak())
}


func main() {
    moncef := Person{"Moncef", 30, 5.7}
    Introduce(moncef)  // This works because Person implements the Speaker interface
}
