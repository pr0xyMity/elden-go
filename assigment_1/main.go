package main

import "fmt"

func main() {
    numbers := []int{}

    for i := 1; i <= 10; i++ {
        numbers = append(numbers, i)
    }

    for n := range numbers {
        if n%2 == 0 {
            fmt.Println(n, "is even") 
        } else {
            fmt.Println(n, "is odd") 
        }
    }
}
