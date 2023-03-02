package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	sum := 0
	lenStudent := len(s.score)

	for _, val := range s.score {
		sum += val
	}

	average := sum / lenStudent

	return float64(average)
}

func (s Student) Min() (min int, name string) {
	minScore := 9999999
	minName := ""

	for index, val := range s.score {
		if minScore > val {
			minScore = val
			minName = s.name[index]
		}
	}

	return minScore, minName
}

func (s Student) Max() (max int, name string) {
	maxScore := 0
	maxName := ""

	for index, val := range s.score {
		if maxScore < val {
			maxScore = val
			maxName = s.name[index]
		}
	}

	return maxScore, maxName
}

func main() {
	var a = Student{}
	for i := 1; i <= 5; i++ {
		var name string
		fmt.Print("Input " + strconv.Itoa(i) + " Student's Name ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\nAverage Score Students is", a.Average())
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is " + nameMin + " (", scoreMin, ")")
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is " + nameMax + " (", scoreMax, ")")
}