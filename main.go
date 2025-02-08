package main

import (
	"management-inventaris/config"
	"management-inventaris/contollers"
	"management-inventaris/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load .env file")
	}

	r := gin.Default()

	db := config.ConnectDatabase()
	db.AutoMigrate(&models.Produk{}, &models.Inventaris{}, &models.Pesanan{})

	productController := contollers.NewProductController(db)

	api := r.Group("/api")
	{
		api.POST("/products", productController.CreateProduct)
		api.GET("/products/:id", productController.GetProduct)
		api.PUT("/products/:id", productController.UpdateProduct)
		api.DELETE("/products/:id", productController.DeleteProduct)
		api.GET("/inventories/:id", productController.GetInventory)
		api.PUT("/inventories/:id", productController.UpdateInventory)
		api.POST("/orders", productController.CreateOrder)
		api.GET("/orders/:id", productController.GetOrder)
	}

	r.Run(":8080")
}
