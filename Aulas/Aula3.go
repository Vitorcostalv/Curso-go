package main

import "fmt"

func main() {
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Println(i, "is even")
        } else {
            fmt.Println(i, "is odd")
        }
    }

    day := "Monday"
    switch day {
    case "Monday", "Wednesday", "Friday":
        fmt.Println("Time to work!")
    case "Saturday", "Sunday":
        fmt.Println("Enjoy the weekend!")
    default:
        fmt.Println("Just another day.")
    }
}
