package main

import "fmt"

type car struct{
	tipe string
	fuel float64
}

func main() {
	var c car

	fmt.Println("Masukan tipe:")
    fmt.Scan(&c.tipe)
	fmt.Println("Masukan jumlah fuel:")
    fmt.Scan(&c.fuel)

    fmt.Println("Kendaraan Anda dengan tipe ", c.tipe, " berjalan sebanyak ", (c.fuel/1.5), " mil")
}