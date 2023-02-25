package main

import "fmt"

func removeDuplicate(array []string) []string {
	var newArray []string
	myMap := make(map[string]bool)
	for _, val := range array {
		if value := myMap[val]; !value {
			myMap[val] = true
			newArray = append(newArray, val)
		}
	}

	return newArray
}

func ArrayMerge(arrayA, arrayB []string) []string {
	var arrayMerge []string
	arrayMerge = append(arrayMerge, arrayA...)
	arrayMerge = append(arrayMerge, arrayB...)

	return removeDuplicate(arrayMerge)
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}