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

func main() {
    moncef := Person{"Moncef", 30, 5.7}
	fmt.Println(moncef.Speak())
    fmt.Println(moncef.Name, "is", moncef.Age, "years old.")
}