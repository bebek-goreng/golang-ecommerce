package db

import "time"

type Address struct {
	Id           int       `gorm:"type:primaryKey" json:"id"`
	IdUser       int       `json:"id_user"`
	JudulAlamat  string    `gorm:"type:varchar(255)" json:"judul_alamat"`
	NamaPenerima string    `gorm:"type:varchar(255)" json:"nama_penerima"`
	NoTlp        string    `gorm:"type:varchar(255)" json:"no_tlp"`
	DetailAlamat string    `gorm:"type:varchar(255)" json:"detail_alamat"`
	UpdatedAT    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}
