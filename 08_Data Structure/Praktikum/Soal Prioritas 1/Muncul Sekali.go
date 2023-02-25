package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int{
	var array []int
	
	for _, val1 := range angka {
		sama := 0
		for _, val2 := range angka {
			if val1 == val2 {
				sama += 1
			}
			
			if sama > 1 {
				break
			}
		}

		if sama == 1 {
			parseInt, _ := strconv.Atoi(string(val1))
			array = append(array, parseInt)
		}
	}

	return array
}

func main(){
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("765723762"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872054"))
}