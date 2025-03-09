package main

import (
	"context"
	"log"
	"management-inventaris/config"
	"management-inventaris/contollers"
	"management-inventaris/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Check Redis connection
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Failed to load .env file")
	}

	initRedis() // Initialize Redis

	r := gin.Default()

	db := config.ConnectDatabase()
	db.AutoMigrate(&models.Produk{}, &models.Inventaris{}, &models.Pesanan{})

	productController := contollers.NewProductController(db, redisClient)
	systemController := contollers.NewSysController(db)

	api := r.Group("/api")
	{
		// Route Untuk CRUD Product
		api.POST("/products", productController.CreateProduct)
		api.GET("/products/:id", productController.GetProduct)
		api.PUT("/products/:id", productController.UpdateProduct)
		api.DELETE("/products/:id", productController.DeleteProduct)
		api.GET("/inventories/:id", productController.GetInventory)
		api.PUT("/inventories/:id", productController.UpdateInventory)
		api.POST("/orders", productController.CreateOrder)
		api.GET("/orders/:id", productController.GetOrder)

		// Route Untuk Download dan Upload File
		api.POST("/upload", systemController.UploadImage)
		api.GET("/download", systemController.DownloadImage)
	}

	r.Run(":8080")
}
