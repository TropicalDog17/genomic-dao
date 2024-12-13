package cmd

import (
	"github.com/TropicalDog17/genomic-dao-service/internal/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db.Connect("host=localhost user=postgres password=postgres dbname=genomic-dao port=5432 sslmode=disable")
}
