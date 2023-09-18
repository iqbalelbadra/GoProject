package models


type Product struct {
    ID_Product     uint   `gorm:"primaryKey autoIncrement" json:"id_produk"`
    NamaProduk    string `gorm:"not null" json:"nama_produk"`
    HargaProduk   float64 `gorm:"not null" json:"harga_produk"`
    GambarProduk  string `json:"gambar_produk"`
    Deskripsi     string `json:"deskripsi_produk"`
    StokProduk    int    `gorm:"not null" json:"stok_produk"`
}

