package main

import (
	"fmt"
	"strings"
)

func caesar(offset int, input string) string {
	chiper := ""
	const alpha = "abcdefghijklmnopqrstuvwxyz"
	const asciiZ = 122
	const totalAlphabet = 26
	const asciiSpace = 32
	
	if offset > totalAlphabet {
		offset %= totalAlphabet
	}

	for _, val := range strings.ToLower(input) {
		index := 0
		index = strings.Index(alpha, string(val)) + 1
		posisi := index + offset 
		if val == asciiSpace {
			chiper += " "
			continue
		}
		
		if val < 97 || val > 122 {
			continue
		}

		if posisi > totalAlphabet {
			chiper += string(rune(asciiZ - totalAlphabet + (posisi - totalAlphabet)))
			continue
		}

		chiper += string(val + rune(offset))
	}

	return chiper
}

func main() {
	fmt.Println(caesar(3, "abc")) // def
	fmt.Println(caesar(2, "alta")) // def
	fmt.Println(caesar(10, "alterraacademy")) // kvdobbkkmknowi 
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz")) // bcdefghijklmnopqrstuvwxyza 
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl 
}