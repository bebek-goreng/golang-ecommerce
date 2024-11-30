package db

import "time"

type Product struct {
	Id            int       `gorm:"primaryKey" json:"id"`
	NamaProduk    string    `gorm:"type:varchar(255); not null" json:"nama_produk"`
	Slug          string    `gorm:"type:varchar(255); not null" json:"slug"`
	HargaReseller string    `gorm:"type:varchar(255); not null" json:"harga_reseller"`
	HargaKonsumen string    `gorm:"type:varchar(255); not null" json:"harga_konsumen"`
	Stock         int       `gorm:"not null" json:"stock"`
	Deskripsi     string    `gorm:"type:varchar(255)" json:"deskripsi"`
	UpdatedAT     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
	CategoryId    int       `json:"category_id"`
	StoreId       int       `json:"store_id"`

	Category Category `gorm:"foreignkey:CategoryId"`
	Store    Store    `gorm:"foreignkey:StoreId"`
}
