package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	products, err := loadProducts("./products.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(products)
}

// Models
type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

// Read products from Memory
func loadProducts(path string) ([]Product, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var list []Product
	err = json.Unmarshal([]byte(file), &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
