package contollers

import (
	"encoding/json"
	"log"
	"management-inventaris/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/redis/go-redis/v9"
)

/*

2. RESTful API Development
    Buat API untuk hal-hal berikut:
    - Produk:
       • Menambahkan, melihat, memperbarui, dan menghapus produk.
       • Melihat detail produk berdasarkan ID atau kategori

    - Inventaris:
       • Melihat tingkat stok untuk suatu produk.
       • Memperbarui tingkat stok (menambah atau mengurangi stok).

    - Pesanan
       • Membuat pesanan baru.
       • Mengambil detail pesanan berdasarkan ID.
*/

type ProductController struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewProductController(db *gorm.DB, redisClient *redis.Client) *ProductController {
	return &ProductController{db: db, redisClient: redisClient}
}

// Menambahkan Produk
func (pc *ProductController) CreateProduct(c *gin.Context) {
	var produk models.Produk
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := pc.db.Create(&produk)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, produk)
}

// Melihat Detail Produk Berdasarkan ID
func (pc *ProductController) GetProduct(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	cacheKey := "product:" + id

	// Check if product exists in Redis
	cachedProduct, err := pc.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		log.Println("Cache hit for product ID:", id)
		var produk models.Produk
		json.Unmarshal([]byte(cachedProduct), &produk)
		c.JSON(200, produk)
		return
	} else if err != redis.Nil {
		log.Println("Redis error:", err)
	}

	// If not in cache, fetch from DB
	var produk models.Produk
	result := pc.db.First(&produk, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}

	// Store product in Redis for 5 minutes
	productJSON, _ := json.Marshal(produk)
	err = pc.redisClient.Set(ctx, cacheKey, productJSON, 5*time.Minute).Err()
	if err != nil {
		log.Println("Failed to cache product:", err)
	}

	c.JSON(200, produk)
}

// Update Produk Berdasarkan ID
func (pc *ProductController) UpdateProduct(c *gin.Context) {
	var produk models.Produk
	id := c.Param("id")
	result := pc.db.First(&produk, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": result.Error.Error()})
		return
	}

	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result = pc.db.Save(&produk)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, produk)
}

// Delete Produk Berdasarkan ID
func (pc *ProductController) DeleteProduct(c *gin.Context) {
	var produk models.Produk
	id := c.Param("id")
	result := pc.db.Delete(&produk, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Product deleted successfully"})
}

// Melihat Tingkat Stok Untuk Suatu Produk
func (pc *ProductController) GetInventory(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	cacheKey := "inventory:" + id

	// Cek apakah data ada di Redis
	cachedInventory, err := pc.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		log.Println("Cache hit for inventory ID:", id)
		var inventaris models.Inventaris
		json.Unmarshal([]byte(cachedInventory), &inventaris)
		c.JSON(200, inventaris)
		return
	} else if err != redis.Nil {
		log.Println("Redis error:", err)
	}

	// Jika tidak ada di cache, ambil dari database
	var inventaris models.Inventaris
	result := pc.db.First(&inventaris, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": result.Error.Error()})
		return
	}

	// Simpan data ke Redis selama 5 menit
	inventoryJSON, _ := json.Marshal(inventaris)
	err = pc.redisClient.Set(ctx, cacheKey, inventoryJSON, 5*time.Minute).Err()
	if err != nil {
		log.Println("Failed to cache inventory:", err)
	}

	c.JSON(200, inventaris)
}

// Memperbarui Tingkat Stok (Menambah atau Mengurangi Stok)
func (pc *ProductController) UpdateInventory(c *gin.Context) {
	var inventaris models.Inventaris
	id := c.Param("id")
	result := pc.db.First(&inventaris, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": result.Error.Error()})
		return
	}

	if err := c.ShouldBindJSON(&inventaris); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result = pc.db.Save(&inventaris)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, inventaris)
}

// Membuat Pesanan Baru
func (pc *ProductController) CreateOrder(c *gin.Context) {
	var pesanan models.Pesanan
	if err := c.ShouldBindJSON(&pesanan); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := pc.db.Create(&pesanan)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(201, pesanan)
}

// Mengambil Detail Pesanan Berdasarkan ID
func (pc *ProductController) GetOrder(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	cacheKey := "order:" + id

	// Cek apakah data ada di Redis
	cachedOrder, err := pc.redisClient.Get(ctx, cacheKey).Result()
	if err == nil {
		log.Println("Cache hit for order ID:", id)
		var pesanan models.Pesanan
		json.Unmarshal([]byte(cachedOrder), &pesanan)
		c.JSON(200, pesanan)
		return
	} else if err != redis.Nil {
		log.Println("Redis error:", err)
	}

	// Jika tidak ada di cache, ambil dari database
	var pesanan models.Pesanan
	result := pc.db.First(&pesanan, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": result.Error.Error()})
		return
	}

	// Simpan data ke Redis selama 5 menit
	orderJSON, _ := json.Marshal(pesanan)
	err = pc.redisClient.Set(ctx, cacheKey, orderJSON, 5*time.Minute).Err()
	if err != nil {
		log.Println("Failed to cache order:", err)
	}

	c.JSON(200, pesanan)
}
