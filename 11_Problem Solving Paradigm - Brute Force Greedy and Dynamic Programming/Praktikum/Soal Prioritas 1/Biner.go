package main

import (
    "fmt"
    "strconv"
)

func main() {
	var number int
    fmt.Printf("Masukkan angka: ")
	fmt.Scanln(&number)

	var arr []string

	for i:=0; i <= number; i++{
		binaryStr := strconv.FormatInt(int64(i), 2)
		arr = append(arr, binaryStr)
	}
    
    fmt.Println(arr)
}