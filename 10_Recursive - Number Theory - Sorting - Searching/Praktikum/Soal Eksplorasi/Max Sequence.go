package main

import "fmt"

func MaxSequence(arr []int) int {
	temp1 := 0
	for i := 0; i < len(arr); i++ {
		myMap := make(map[int]bool)
		temp2 := int(arr[i])
		for j := i + 1; j < len(arr); j++ {
			temp2 += int(arr[j])
			myMap[temp2] = true
		}

		for k := range myMap {
			if k > temp1 {
				temp1 = k
			}
		}
	}

	return temp1

}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}