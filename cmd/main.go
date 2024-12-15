package main

import (
	_ "github.com/TropicalDog17/genomic-dao-service/api"
	"github.com/TropicalDog17/genomic-dao-service/internal/server"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	server.StartService(db)

}
