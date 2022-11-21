package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Inventory   string  `json:"inventory"`
	Code        string  `json:"code"`
	IsPublished bool    `json:"isPublished"`
	CreatedAt   string  `json:"createdAt"`
}

func main() {
	app := gin.Default()

	res, err := os.ReadFile("./products.json")
	if err != nil {
		panic("Could not read file.")
	}

	var products []Product

	err = json.Unmarshal(res, &products)
	if err != nil {
		panic(err)
	}

	app.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, products)
		fmt.Println(products)
	})

	app.Run(":8080")
}
