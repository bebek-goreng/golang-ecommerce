package migration

import (
	"fmt"
	"log"

	"github.com/bebek-goreng/golang-ecommerce/internal/models/db"
	"gorm.io/gorm"
)

func AutoMigrate(database *gorm.DB) {
	err := database.AutoMigrate(
		&db.User{},
		&db.Category{},
		&db.Product{},
		&db.Store{},
		&db.Transaction{},
		&db.Address{},
		&db.DetailTrx{},
		&db.LogProduct{},
	)

	if err != nil {
		log.Fatal("Migration db failed:", err)
	}

	fmt.Println("Migration success")
}
