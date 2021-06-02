package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Product struct {
	Name     string
	Category string
	Price    int
}

func randomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func main() {

	names := []string{}
	for i := 0; i < 50; i++ {
		productName := "Product " + strconv.Itoa(i)
		names = append(names, productName)
	}
	category := [4]string{"fashion", "electronics", "sport", "food"}
	products := [20]Product{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(products); i++ {
		products[i] = Product{
			names[rand.Intn(len(names))],
			category[rand.Intn(len(category))],
			randomInt(100, 200),
		}
	}
	fmt.Println("All products: ")
	for _, product := range products {
		fmt.Println(product)
	}

	fmt.Println("Search product by name: ")
	var searchKey string
	fmt.Print("Key name = ")
	fmt.Scan(&searchKey)
	for i := 0; i < len(products); i++ {
		if strings.Contains(products[i].Name, searchKey) {
			fmt.Println(products[i])
		}
	}

	fmt.Println("Get product by category: ")
	var categoryStr string
	fmt.Print("Category = ")
	fmt.Scan(&categoryStr)
	for i := 0; i < len(products); i++ {
		if products[i].Category == categoryStr {
			fmt.Println(products[i])
		}
	}

	fmt.Println("Get product by price")
	var minPrice int
	fmt.Print("Min Price = ")
	fmt.Scan(&minPrice)

	var maxPrice int
	fmt.Print("Max Price = ")
	fmt.Scan(&maxPrice)

	for i := 0; i < len(products); i++ {
		if products[i].Price <= maxPrice && products[i].Price >= minPrice {
			fmt.Println(products[i])
		}
	}
}
