package main

import (
    "fmt"
)

func main() {
    var A, B, C int
    fmt.Println("Enter values for A, B, and C:")
    fmt.Scan(&A, &B, &C)

	var arr []int

    for x := 1; x <= 100; x++ {
        for y := 1; y <= 100; y++ {
            for z := 1; z <= 100; z++ {
                if x+y+z == A && x*y*z == B && x*x+y*y+z*z == C {
                    arr = append(arr, x)
					arr = append(arr, y)
					arr = append(arr, z)					
				} 
				if len(arr) != 0{
					break
				}
			}
			if len(arr) != 0{
				break
			}
		}
		if len(arr) != 0{
			break
		}
	}
	if len(arr) == 0{
		fmt.Println("no solution")
	} else {
		fmt.Println(arr[0], arr[1], arr[2])
	}
}

