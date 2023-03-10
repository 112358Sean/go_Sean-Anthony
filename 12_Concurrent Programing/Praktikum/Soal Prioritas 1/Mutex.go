package main

import (
    "fmt"
    "sync"
)

var mutex = &sync.Mutex{}

func factorial(n int, wg *sync.WaitGroup) {
    defer wg.Done()
    mutex.Lock()
    defer mutex.Unlock()
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    fmt.Println("Faktorial dari" , n , "adalah" , result)
}

func main() {
    var wg sync.WaitGroup
    for i := 1; i <= 5; i ++{
        wg.Add(1)
        go factorial(i, &wg)
    }
    wg.Wait()
}