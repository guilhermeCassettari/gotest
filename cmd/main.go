package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	usecase "go-api/useCase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// reflex -r '\.go$' -- go run main.go

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductRepository := repository.NewProductRepository(dbConnection)
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.POST("/product", ProductController.CreateProduct)

	server.GET("/product/:id", ProductController.GetProductById)

	server.Run(":8080")
}
