package main

import "fmt"

func Pairtarget(arr []int, target int) []int {
	var numbers []int
	for i1, val1 := range arr {
		for i2, val2 := range arr {
			if val1+val2 == target {
				numbers = append(numbers, i1)
				numbers = append(numbers, i2)
				break
			}
		}
		if len(numbers) != 0 {
			break
		}
	}
	return numbers
}

func main() {
	fmt.Println(Pairtarget([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(Pairtarget([]int{2, 5, 9, 11}, 11))
	fmt.Println(Pairtarget([]int{1, 3, 5, 7}, 12))
	fmt.Println(Pairtarget([]int{1, 4, 6, 8}, 10))
	fmt.Println(Pairtarget([]int{1, 5, 6, 7}, 6))
}