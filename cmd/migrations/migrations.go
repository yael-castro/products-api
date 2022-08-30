package main

import (
	"github.com/yael-castro/products-api/internal/model"
	"github.com/yael-castro/products-api/internal/repository"
	"log"
	"os"
)

func main() {
	db, err := repository.NewGormDB(os.Getenv("GORM_DSN"))
	if err != nil {
		log.Fatal(err)
	}

	db.Migrator().CreateTable(&model.Product{})
}
