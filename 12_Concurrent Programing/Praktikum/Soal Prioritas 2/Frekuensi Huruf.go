package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	frekuensi := make(map[string]int)
	inputan := "lorem ipsum dolor sit amet"

	wg.Add(1)

	go func() {
		for _, inputan := range inputan[:] {
			frekuensi[string(inputan)] += 1
		}
		for key, val := range frekuensi {
			fmt.Println(key, ":", val)
		}
		defer wg.Done()
	}()

	wg.Wait()
}