package main

import (
    "fmt"
    "strings"
)

func Mapping(slice []string)map[string]int{
    mapped := make(map[string]int)
    for _, word := range slice {
        word = strings.ToLower(word)
        if _, ok := mapped[word]; ok {
            mapped[word]++
        } else {
            mapped[word] = 1
        }
    }

    return mapped
}

func main(){
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}