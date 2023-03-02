package main

import (
	"fmt"
)

type pair struct {
	name  string
	count int
}

func BubbleSort(pairs []pair) {
	for i := 0; i < len(pairs)-1; i++ {
		for j := 0; j < len(pairs)-i-1; j++ {
			if pairs[j].count > pairs[j+1].count {
				pairs[j], pairs[j+1] = pairs[j+1], pairs[j]
			}
		}
	}
}

func MostAppearItem(items []string) []pair {
	var pairs []pair
	myMap := make(map[string]int)

	for _, val := range items {
		myMap[val] += 1
	}

	for key, val := range myMap {
		var p pair
		p.name = key
		p.count = val
		pairs = append(pairs, p)
	}

	BubbleSort(pairs)

	return pairs
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}