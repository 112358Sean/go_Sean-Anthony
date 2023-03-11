package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Product struct {
    Title    string  `json:"title"`
    Price    float64 `json:"price"`
    Category string  `json:"category"`
}

func main() {
    productsChan := make(chan []Product)

    go func() {
        resp, err := http.Get("https://fakestoreapi.com/products")
        if err != nil {
            panic(err)
        }
        defer resp.Body.Close()

        var products []Product
        err = json.NewDecoder(resp.Body).Decode(&products)
        if err != nil {
            panic(err)
        }

        productsChan <- products
    }()

    products := <-productsChan

	fmt.Println("products data")
    for _, p := range products {
		fmt.Println("===")
        fmt.Println("Title:", p.Title)
		fmt.Println("Price:", p.Price)
		fmt.Println("Category:", p.Category)
		fmt.Println("===")
    }
}