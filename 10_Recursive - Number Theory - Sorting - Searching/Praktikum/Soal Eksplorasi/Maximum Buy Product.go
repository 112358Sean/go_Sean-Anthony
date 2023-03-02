package main

import (
	"fmt"
)

func BubbleSort(array []int) {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func MaximumBuyProduct(money int, productPrice []int) {
	if money == 0 {
		fmt.Println(0)
		return
	}
	
	sum := 0
	for _, val := range productPrice {
		sum += val
	}

	selisih := sum - money
	BubbleSort(productPrice)

	dibuang := 0
	for i := 0; i <= len(productPrice) - 1; i++ {
		if selisih > 0 {
			selisih -= int(productPrice[i])
			dibuang += 1
		}

		if selisih < 0 {
			break
		}
	}

	fmt.Println(len(productPrice) - dibuang)
}

func main() {
	MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000})
	MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000})
	MaximumBuyProduct(10000, []int{2000, 3000, 1000, 2000, 10000})
	MaximumBuyProduct(4000, []int{7500, 3000, 2500, 2000})
	MaximumBuyProduct(0, []int{10000, 30000})
}