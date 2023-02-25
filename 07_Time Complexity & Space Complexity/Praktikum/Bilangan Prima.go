package main

import (
	"fmt"
	"math"
)

func isPrime(num int) bool {
	if num == 1 {
		return false
	}

	for i := 3.0; i < math.Sqrt(float64(num)); i++ {
		if num % int(i) == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPrime(1000000007))
	fmt.Println(isPrime(1500450271))
	fmt.Println(isPrime(13))
	fmt.Println(isPrime(15))
	fmt.Println(isPrime(20))
}