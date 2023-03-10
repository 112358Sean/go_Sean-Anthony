package main

import (
    "fmt"
    "time"
)

func printMultiples(x int) {
    for i := 1; i <= 5; i++ {
        fmt.Println(i*x)
		time.Sleep(1 * time.Second)
    }
}

func main() {
    go printMultiples(2)
    go printMultiples(3)
    go printMultiples(4)

    time.Sleep(11 * time.Second)
}