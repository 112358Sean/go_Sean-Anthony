package main

import "fmt"

func main(){
	var number int
    fmt.Printf("Masukkan angka: ")
	fmt.Scanln(&number)

    roman := ""
    values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    for i := 0; i < len(values); i++ {
        for number >= values[i] {
            number -= values[i]
            roman += symbols[i]
        }
    }
    fmt.Println(roman)
}