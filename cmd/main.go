package main

import (
	"log"
	"medical-record-api/config"
	"medical-record-api/internal/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load .env file:" + err.Error())
	}

	// init db and routes
	db := config.NewDB()
	r := router.SetupRouter(db)

	// run app
	port := os.Getenv("PORT")
	r.Run("0.0.0.0:" + port)
}
