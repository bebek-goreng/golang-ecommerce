package db

import "time"

type DetailTrx struct {
	Id           int       `gorm:"primaryKey" json:"id"`
	TrxId        int       `json:"id_trx"`
	IdLogProduct int       `json:"id_log_product"`
	IdToko       int       `json:"id_toko"`
	Kuantitas    int       `gorm:"not null" json:"kuantitas"`
	HargaTotal   int       `gorm:"not null" json:"harga_total"`
	UpdatedAT    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`

	Transaction Transaction `gorm:"foreignkey:TrxId"`
	LogProduct  LogProduct  `gorm:"foreignkey:IdLogProduct"`
}
