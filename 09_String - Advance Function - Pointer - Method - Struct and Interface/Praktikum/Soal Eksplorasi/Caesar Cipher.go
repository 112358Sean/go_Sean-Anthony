package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	nameEncode string
	score      int
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *student) Encode() string {
	var nameEncode = ""
	letter := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	dict := make(map[string]string)
	key := s.score

	for i := range letter {
		dict[string(letter[i])] = string(letter[(int(i)+key)%len(letter)])
	}

	var cipher_txt []string

	for i := range s.name {
		if strings.Contains(letter, string(s.name[i])) {
			temp := dict[string(s.name[i])]
			cipher_txt = append(cipher_txt, temp)
		} else {
			temp := string(s.name[i])
			cipher_txt = append(cipher_txt, temp)
		}
	}

	nameEncode = strings.Join(cipher_txt, "")

	return nameEncode
}

func (s *student) Decode() string {
	var nameDecode = ""

	letter := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	dict := make(map[string]string)
	key := s.score

	for i := range letter {
		if int(i)-key < 0 {
			dict[string(letter[i])] = string(letter[(26-int(i)-key)%len(letter)])
			continue
		}
		dict[string(letter[i])] = string(letter[(int(i)-key)%len(letter)])
	}

	var plain_txt []string

	for i := range s.nameEncode {
		if strings.Contains(letter, string(s.nameEncode[i])) {
			temp := dict[string(s.nameEncode[i])]
			plain_txt = append(plain_txt, temp)
		} else {
			temp := string(s.nameEncode[i])
			plain_txt = append(plain_txt, temp)
		}
	}

	nameDecode = strings.Join(plain_txt, "")

	return nameDecode
}

func main() {
	var menu int
	var a = student{}
	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ? ")
	fmt.Scan(&menu)
	if menu == 1 {
		fmt.Print("\nInput Student's Name : ")
		fmt.Scan(&a.name)
		a.score = 4
		fmt.Print("\nEncode of Student's Name " + a.name + " is : " + c.Encode())
	} else if menu == 2 {
		fmt.Print("\nInput Student's Encode Name : ")
		fmt.Scan(&a.nameEncode)
		a.score = 4
		fmt.Print("\nDecode of Student's Name " + a.nameEncode + " is : " + c.Decode())
	} else {
		fmt.Println("Wrong input name menu")
	}
}