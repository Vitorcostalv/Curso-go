package main

import "fmt"

type Animal interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof! Woof!"
}

func main() {
    var animal Animal
    animal = Dog{Name: "Buddy"}
    fmt.Println(animal.Speak())
}
