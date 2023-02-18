package main

import "fmt"

func main(){
	var nilai int
	fmt.Printf("Masukan Nilai: ")
	fmt.Scanln(&nilai)

	if (100 >= nilai && nilai >= 80){
		fmt.Println("Grade : A")
	} else if(79 >= nilai && nilai >= 65){
		fmt.Println("Grade : B")
	} else if(64 >= nilai && nilai >= 50){
		fmt.Println("Grade : C")
	} else if(49 >= nilai && nilai >= 35){
		fmt.Println("Grade : D")
	} else if(34 >= nilai && nilai >= 0){
		fmt.Println("Grade : E")
	} else{
		fmt.Println("Nilai Invalid")
	}
}