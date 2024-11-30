package db

import "time"

type Transaction struct {
	Id               int       `gorm:"primaryKey" json:"id"`
	IdUser           int       `json:"id_user" gorm:"not null"`
	AlamatPengiriman int       `json:"alamat_pengiriman" gorm:"not null"`
	HargaTotal       int       `json:"harga_total" gorm:"not null"`
	KodeInvoice      string    `gorm:"type:varchar(255);not null" json:"kode_invoice"`
	MethodBayar      string    `gorm:"type:varchar(255);not null" json:"method_bayar"`
	UpdatedAT        time.Time `json:"updated_at"`
	CreatedAt        time.Time `json:"created_at"`

	User    User    `gorm:"foreignkey:IdUser"`
	Address Address `gorm:"foreignkey:AlamatPengiriman"`
}
