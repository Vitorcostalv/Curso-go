package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    var numbers = [5]int{1, 2, 3, 4, 5}
    fmt.Println("Array:", numbers)

    fruits := []string{"Apple", "Banana", "Orange"}
    fmt.Println("Slice:", fruits)

    ages := map[string]int{
        "Alice": 30,
        "Bob":   25,
    }
    fmt.Println("Map:", ages)

    person := Person{Name: "John", Age: 28}
    fmt.Println("Struct:", person)
}
