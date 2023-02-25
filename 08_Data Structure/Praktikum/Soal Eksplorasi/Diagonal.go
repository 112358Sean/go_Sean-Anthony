package main

import "fmt"

func main() {
	var rows, cols int

    fmt.Println("Masukan sisi : ")
    fmt.Scan(&rows)

	cols = rows

    arr := make([][]int, rows)
    for i := range arr {
        arr[i] = make([]int, cols)
    }

    fmt.Println("Masukan angka: ")
    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            fmt.Scan(&arr[i][j])
        }
    }

    fmt.Println("Array 2D adalah:", arr)

    count1 := 0
    count2 := 0

    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr[i]); j++ {
            if i == j {
                count1 += arr[i][j]
            }
            if i+j == len(arr)-1 {
                count2 += arr[i][j]
            }
        }
    }

    if count1 > count2{
		fmt.Println("Hasil nilai mutlak pengurangan diagonal adalah", count1-count2)
	} else {
		fmt.Println("Hasil nilai mutlak pengurangan diagonal adalah", count2-count1)
	}
}
