package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id           int
	Nama         string
	NoTlp        string
	TanggalLahir time.Time
	JenisKelamin string
	Tentang      string
	Pekerjaan    string
	Email        string
	IdProvinsi   string
	IdKota       string
	IsAdmin      bool
	Address      []Address
}
