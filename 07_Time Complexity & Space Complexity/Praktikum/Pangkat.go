package main

import "fmt"

func Pangkat(angka, pangkat int) int {
	if pangkat == 0 {
		return 1
	}

	tmp := Pangkat(angka, pangkat/2)

	if pangkat % 2 == 0 {
		return tmp * tmp
	}

	if pangkat > 0 {
		return tmp * tmp * angka
	}

	return tmp * tmp / angka
}

func main() {
	fmt.Println(Pangkat(2, 3))
	fmt.Println(Pangkat(5, 3))
	fmt.Println(Pangkat(10, 2))
	fmt.Println(Pangkat(2, 5))
	fmt.Println(Pangkat(7, 3))
}