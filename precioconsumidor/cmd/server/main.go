package main

import (
	"github.com/bootcamp-go/Consignas-Go-Web.git/cmd/server/handler"
	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/product"
	"github.com/bootcamp-go/Consignas-Go-Web.git/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {

	/* 	if err := godotenv.Load("./cmd/server/.env"); err != nil {
	   		panic("Error loading .env file: " + err.Error())
	   	}
	*/
	storage := store.NewStore("../../products.json")

	repo := product.NewRepository(storage)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", productHandler.GetAll())
		products.GET(":id", productHandler.GetByID())
		products.GET("/search", productHandler.Search())
		products.POST("", productHandler.Post())
		products.DELETE(":id", productHandler.Delete())
		products.PATCH(":id", productHandler.Patch())
		products.PUT(":id", productHandler.Put())
		products.GET("/consumer_price", productHandler.ConsumerPrice())
	}
	r.Run(":8080")
}