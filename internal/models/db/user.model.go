package db

import "time"

type User struct {
	Id           int       `gorm:"primaryKey" json:"id"`
	Nama         string    `gorm:"type:varchar(255); not null" json:"nama"`
	KataSandi    string    `gorm:"type:varchar(255); not null" json:"kata_sandi"`
	NoTlp        string    `gorm:"type:varchar(255); unique" json:"no_tlp"`
	TanggalLahir time.Time `gorm:"type:date" json:"tanggal_lahir"`
	JenisKelamin string    `gorm:"type:varchar(255)" json:"jenis_kelamin"`
	Tentang      string    `gorm:"type:text" json:"tentang"`
	Pekerjaan    string    `gorm:"type:varchar" json:"pekerjaan"`
	Email        string    `gorm:"type:varchar(255); not null;unique" json:"email"`
	IdProvinsi   string    `gorm:"type:varchar(255)" json:"id_provinsi"`
	IdKota       string    `gorm:"type:varchar(255)" json:"id_kota"`
	IsAdmin      bool      `gorm:"type:boolean" json:"is_admin"`
	UpdatedAT    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`

	Address []Address `gorm:"foreignkey:Id"`
}
