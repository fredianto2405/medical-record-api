package main

import (
	"log"
	"medical-record-api/config"
	"medical-record-api/internal/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load env file")
	}

	db := config.NewDB()
	r := router.SetupRouter(db)
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
