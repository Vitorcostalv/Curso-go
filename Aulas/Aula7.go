package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(500 * time.Millisecond)
    }
}

func printLetters() {
    for i := 'A'; i <= 'E'; i++ {
        fmt.Printf("%c\n", i)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    go printNumbers()
    go printLetters()

    time.Sleep(3 * time.Second)
    fmt.Println("Done!")
}
