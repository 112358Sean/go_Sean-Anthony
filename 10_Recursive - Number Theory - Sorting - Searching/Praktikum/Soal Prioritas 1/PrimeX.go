package main

import (
	"fmt"
)

func checkPrime(num int) bool {
	if num == 1 {
		return false
	}

	for i := 2; i < num; i++ {
		if num % i == 0 {
			return false
		}
	}

	return true
}

func primeX(number int) int {
	var output int
	var index int
	count := 2

	for {
		result := checkPrime(count)
		if result {
			index += 1
		}

		if index == number {
			output = count
			break
		}

		count += 1
	}

	return output
}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}