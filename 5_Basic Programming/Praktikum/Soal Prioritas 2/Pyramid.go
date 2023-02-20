package main

import "fmt"

func main() {
    var baris int
    fmt.Printf("Tentukan jumlah baris: ")
    fmt.Scanln(&baris)

    for i := 1; i <= baris; i++ {
        for j := 1; j <= baris-i; j++ {
            fmt.Printf(" ")
        }
        for k := 1; k <= i; k++ {
            fmt.Printf("* ")
        }
        fmt.Println("")
    }
}