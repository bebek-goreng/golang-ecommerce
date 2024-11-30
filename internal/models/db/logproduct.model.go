package db

import "time"

type LogProduct struct {
	Id            int       `gorm:"primaryKey" json:"id"`
	IdProduk      int       `gorm:"not null" json:"id_produk"`
	NamaProduk    string    `gorm:"type:varchar(255);not null" json:"nama_produk"`
	Slug          string    `gorm:"type:varchar(255);not null" json:"slug"`
	HargaReseller string    `gorm:"type:varchar(255); not null" json:"harga_reseller"`
	HargaKonsumen string    `gorm:"type:varchar(255); not null" json:"harga_konsumen"`
	Deskripsi     string    `gorm:"type:varchar(255)" json:"deskripsi"`
	UpdatedAT     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
	IdToko        int       `gorm:"not null" json:"id_toko"`
	IdCategory    int       `gorm:"not null" json:"id_category"`

	Product  Product  `gorm:"foreignkey:IdProduk"`
	Category Category `gorm:"foreignkey:IdCategory"`
}
