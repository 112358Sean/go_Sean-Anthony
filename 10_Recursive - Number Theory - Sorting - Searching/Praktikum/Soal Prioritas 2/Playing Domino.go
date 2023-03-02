package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	for _, val1 := range cards {
		for _, val2 := range deck {
			if val2 == int(val1[0]) || val2 == int(val1[1]) {
				return val1
			}
		}
	}

	return "tutup kartu"
}

func main() {
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))
}