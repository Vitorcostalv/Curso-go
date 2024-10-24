package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    sum := add(5, 3)
    fmt.Println("Sum:", sum)

    // Função anônima
    multiply := func(x int, y int) int {
        return x * y
    }
    fmt.Println("Multiplication:", multiply(4, 2))
}
