package db

import "time"

type Store struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	UserId    int       `gorm:"not null" json:"user_id"`
	NamaToko  string    `gorm:"type:varchar(255);not null" json:"nama_toko"`
	UrlFoto   string    `gorm:"type:varchar(255)" json:"url_foto"`
	UpdatedAT time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`

	User       User         `gorm:"foreignkey:UserId"`
	Product    []Product    `gorm:"foreignkey:StoreId"`
	DetailTrx  []DetailTrx  `gorm:"foreignkey:IdToko"`
	LogProduct []LogProduct `gorm:"foreignkey:IdToko"`
}
