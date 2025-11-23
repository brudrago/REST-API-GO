package main

import (
	"database/sql"
	"log"
	"time"

	"rest-api/controller"
	"rest-api/db"
	"rest-api/repository"
	"rest-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	// === TENTA CONECTAR NO BANCO COM RETRY ===
	var dbConnection *sql.DB
	var err error

	for i := 1; i <= 10; i++ {
		dbConnection, err = db.ConnectDB()
		if err == nil {
			break
		}

		log.Printf("DB não está pronto ainda (tentativa %d/10): %v", i, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Não consegui conectar no banco depois de várias tentativas: %v", err)
	}
	// =========================================

	productRepository := repository.NewProductRepository(dbConnection)
	productUseCase := usecase.NewProductUseCase(productRepository)
	productController := controller.NewProductController(productUseCase)

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	server.GET("/products", productController.GetProducts)
	server.POST("/products", productController.CreateProduct)
	server.PUT("/products/:productId", productController.UpdateProduct)
	server.DELETE("/products/:productId", productController.DeleteProduct)

	if err := server.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
