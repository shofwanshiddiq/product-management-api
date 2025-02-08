package contollers

import (
	"management-inventaris/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	db *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{db: db}
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
	var produk models.Produk
	id := c.Param("id")
	result := pc.db.First(&produk, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": result.Error.Error()})
		return
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
	var inventaris models.Inventaris
	id := c.Param("id")
	result := pc.db.First(&inventaris, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": result.Error.Error()})
		return
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
	var pesanan models.Pesanan
	id := c.Param("id")
	result := pc.db.First(&pesanan, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, pesanan)
}
