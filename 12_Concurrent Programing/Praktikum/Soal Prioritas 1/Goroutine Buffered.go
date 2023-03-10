package main

import (
	"fmt"
)

func main() {
	results := make(chan int, 5)

	go func() {
		for i := 0; i <= 5; i++ {
			results <- i * 3
		}
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}
}