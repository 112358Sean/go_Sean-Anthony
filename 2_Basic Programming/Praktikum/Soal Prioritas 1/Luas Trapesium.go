package main

import "fmt"

func main(){
	var Tinggi int
	fmt.Printf("Masukan Tinggi: ")
	fmt.Scanln(&Tinggi)
	var Sisi_Atas int
	fmt.Printf("Masukan Sisi Atas: ")
	fmt.Scanln(&Sisi_Atas)
	var Sisi_Bawah int
	fmt.Printf("Masukan Sisi Bawah: ")
	fmt.Scanln(&Sisi_Bawah)

	Luas := (Sisi_Atas+Sisi_Bawah)*Tinggi/2
	fmt.Println("Luas Trapesium adalah ", Luas)
}