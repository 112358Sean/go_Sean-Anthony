package main

import (
    "bufio"
    "fmt"
    "os"
)

func main(){
	reader := bufio.NewScanner(os.Stdin)
    fmt.Print("Masukan kata/ kalimat: ")
    reader.Scan()
    kata := reader.Text()

	for i := 0; i < len(kata)-1; i++{
		if kata[i] != kata[len(kata)-1-i]{
			fmt.Println("Bukan Palindrome")
			break
		} else {
			fmt.Println("Palindrome")
			break
		}
	}
}