package main

import (
	"log"
	"os"

	"github.com/bebek-goreng/golang-ecommerce/config"
	"github.com/bebek-goreng/golang-ecommerce/migration"
	"github.com/bebek-goreng/golang-ecommerce/pkg/utils"
	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadEnv()
	db := config.ConnectDb()
	migration.AutoMigrate(db)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("Port is not set in env file")
	}

	r := gin.Default()

	r.Run(":" + port)
}
