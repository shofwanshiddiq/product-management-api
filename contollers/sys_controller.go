package contollers

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/*
File Handling and Storage
     Tambahan utama dalam tugas ini adalah unggah dan unduh gambar produk, dengan spesifikasi berikut:

     - Endpoint Upload: Mengunggah gambar produk ke penyimpanan lokal

     - Endpoint Download: Mengunduh gambar produk berdasarkan ID produk.

     - Penyimpanan: Bisa menggunakan folder lokal

     - Validasi File: Pastikan hanya format gambar tertentu yang diterima (misal: PNG, JPG, JPEG) dengan ukuran maksimum yang ditentukan.

     - Error Handling: Tangani skenario kesalahan seperti file berukuran terlalu besar atau format yang tidak sesuai.
*/

type SysController struct {
	DB            *gorm.DB
	downloadMutex *sync.Mutex
}

func NewSysController(db *gorm.DB) *SysController {
	return &SysController{
		DB:            db,
		downloadMutex: &sync.Mutex{},
	}
}

//Endpoint Upload: Mengunggah gambar produk ke penyimpanan lokal berdasarkan id produk

func (sc *SysController) UploadImage(c *gin.Context) {
	// Ambil nama file id
	fileId := c.DefaultPostForm("file_id", "")
	if fileId == "" {
		c.JSON(400, gin.H{"error": "file_id is required"})
		return
	}

	// Ambil file dari file data
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to read file"})
		return
	}
	defer file.Close()

	// Validasi content image
	contentType := fileHeader.Header.Get("Content-Type")
	allowedContentTypes := map[string]string{
		"image/png":  ".png",
		"image/jpeg": ".jpg",
	}

	fileExt, exists := allowedContentTypes[contentType]
	if !exists {
		c.JSON(400, gin.H{"error": "Invalid file format. Only PNG, JPG, and JPEG are allowed."})
		return
	}

	// Validasi file size
	const maxFileSize = 10 * 1024 * 1024 // 10 MB
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to read file content"})
		return
	}

	if len(fileData) > maxFileSize {
		c.JSON(400, gin.H{"error": "File size exceeds the maximum allowed size of 10 MB."})
		return
	}

	filePath := filepath.Join("upload", fileId+fileExt)
	outFile, err := os.Create(filePath)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to save file"})
		return
	}
	defer outFile.Close()

	_, err = outFile.Write(fileData)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to write file"})
		return
	}

	c.JSON(200, gin.H{"message": "File uploaded successfully", "file_path": filePath})
}

func (sc *SysController) DownloadImage(c *gin.Context) {
	sc.downloadMutex.Lock()
	defer sc.downloadMutex.Unlock()

	fileId := c.Query("file_id")
	filePath := filepath.Join("upload", fileId)

	if strings.TrimSpace(fileId) == "" {
		c.JSON(400, gin.H{"error": "File ID is required"})
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %v", err) // Log the error for debugging
		c.JSON(404, gin.H{"error": "File not found"})
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Printf("Error getting file info: %v", err) // Log the error for debugging
		c.JSON(500, gin.H{"error": "Failed to get file info"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+fileId)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

	done := make(chan bool)
	errChan := make(chan error)

	go func() {
		defer close(done)
		buffer := make([]byte, 4096)

		for {
			n, err := file.Read(buffer)
			if n > 0 {
				_, writeErr := c.Writer.Write(buffer[:n])
				if writeErr != nil {
					errChan <- fmt.Errorf("failed to write file data")
					return
				}
				c.Writer.Flush()
			}

			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				errChan <- fmt.Errorf("failed to read file")
				return
			}
		}
	}()

	select {
	case <-done:
		return
	case err := <-errChan:
		log.Printf("Error during file read: %v", err) // Log the error for debugging
		c.JSON(500, gin.H{"error": err.Error()})
	case <-time.After(5 * time.Second):
		c.JSON(http.StatusRequestTimeout, gin.H{"error": "Request timeout"})
	}
}
