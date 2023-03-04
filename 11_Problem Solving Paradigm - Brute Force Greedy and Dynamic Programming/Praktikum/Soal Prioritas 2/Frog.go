package main

import(
	"fmt"
)

func Frog(h []int) int {
    cost := make([]int, len(h))
    for i := range cost {
        cost[i] = 0
    }

    if len(h) == 1 {
        return cost[0]
    }

    if len(h) == 2 {
        return abs(h[1] - h[0])
    }

    cost[0] = 0

    cost[1] = abs(h[1] - h[0])

    for i := 2; i < len(h); i++ {
        cost[i] = min(cost[i-1]+abs(h[i]-h[i-1]), cost[i-2]+abs(h[i]-h[i-2]))
    }

    return cost[len(h)-1]
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
    fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}
