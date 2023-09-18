package models


type Transaksi struct {
    ID_Transaksi      uint   `gorm:"primaryKey autoIncrement" json:"id_transaksi"`
    ID_User           uint   `gorm:"not null" json:"id_user"`
    ID_Product        uint   `gorm:"not null" json:"id_produk"`
    TanggalTransaksi string `gorm:"not null" json:"tanggal_transaksi"`
    JumlahBarang     int    `gorm:"not null" json:"jumlah_barang"`
    StatusPembayaran string `gorm:"not null" json:"status_pembayaran"`
    MetodePembayaran string `gorm:"not null" json:"metode_pembayaran"`
}
