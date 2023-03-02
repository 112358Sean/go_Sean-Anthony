package main

import (
	"fmt"
)

func Compare(a, b string) string {
	common := ""
	for index1, value1 := range a {
		for index2, value2 := range b {
			if string(value1) == string(value2) && index1 == index2 {
				common += string(value2)
			}
		}
	}
	return common
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}