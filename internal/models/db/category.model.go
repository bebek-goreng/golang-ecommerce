package db

import "time"

type Category struct {
	Id           int       `gorm:"primaryKey" json:"id"`
	NamaCategory string    `gorm:"type:varchar(255)" json:"nama_kategori"`
	UpdatedAT    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
}
