package main

import (
	"os"

	"github.com/bootcamp-go/Consignas-Go-Web/cmd/server/handler"
	"github.com/bootcamp-go/Consignas-Go-Web/internal/product"
	"github.com/bootcamp-go/Consignas-Go-Web/pkg/middleware"
	"github.com/bootcamp-go/Consignas-Go-Web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/example/gzipped/docs"
)

// @title Store API
// @version 1.0
// @description This API handle products from our store
// @termsOfService https://publicapis.ml/terms

// @contact.name Public APIs
// @contact.url https://dev.publicapis.ml

// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html


func main() {

	if err := godotenv.Load("./.env"); err != nil {
		panic("Error loading .env file: " + err.Error())
	}

	storage := store.NewStore("./products.json")

	repo := product.NewRepository(storage)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())

	docs.SwaggerInfo_swagger.Host = os.Getenv("HOST")
   	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET("", productHandler.GetAll())
		products.GET(":id", productHandler.GetByID())
		products.GET("/search", productHandler.Search())
		products.GET("/consumer_price", productHandler.ConsumerPrice())
		products.POST("", middleware.Authentication(), productHandler.Post())
		products.DELETE(":id", middleware.Authentication(), productHandler.Delete())
		products.PATCH(":id", middleware.Authentication(), productHandler.Patch())
		products.PUT(":id", middleware.Authentication(), productHandler.Put())
	}
	r.Run(":8080")
}
