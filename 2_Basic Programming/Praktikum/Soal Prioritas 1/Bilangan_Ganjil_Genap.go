package main

import "fmt"

func main(){
	var number int
	fmt.Printf("Masukan Angka: ")
	fmt.Scanln(&number)

	if (number % 2 == 0){
		fmt.Println("Bilangan Genap")
	} else{
		fmt.Println("Bilangan Ganjil")
	}
}