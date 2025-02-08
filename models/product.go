package models

import (
	"gorm.io/gorm"
)

/*
Database Design and Queries
a. Desain sebuah database relasional dengan setidaknya tabel-tabel berikut:
     - Produk: Menyimpan detail produk (ID, nama, deskripsi, harga, dan kategori).
     - Inventaris: Melacak tingkat stok dan lokasi produk (ID produk, jumlah, dan lokasi)
     - Pesanan: Mencatat pesanan pelanggan (ID pesanan, ID produk, jumlah, dan tanggal
*/

type Produk struct {
	gorm.Model
	ID        uint   `json:"id"`
	Nama      string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
	Harga     int    `json:"harga"`
	Kategori  string `json:"kategori"`
}

type Inventaris struct {
	gorm.Model
	IDProduk uint   `json:"id_produk"`
	Jumlah   int    `json:"jumlah"`
	Lokasi   string `json:"lokasi"`
}

type Pesanan struct {
	gorm.Model
	IDPesanan uint   `json:"id_pesanan"`
	IDProduk  uint   `json:"id_produk"`
	Jumlah    int    `json:"jumlah"`
	Tanggal   string `json:"tanggal"`
}
