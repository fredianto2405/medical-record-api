package main

import (
	"log"
	"medical-record-api/config"
	"medical-record-api/internal/email"
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

	// email scheduler
	emailRepo := email.NewRepository(db)
	emailService := email.NewService(emailRepo)
	emailCfg := email.Config{
		SMTPHost: os.Getenv("MAIL_SMTP_HOST"),
		SMTPPort: os.Getenv("MAIL_SMTP_PORT"),
		Sender:   os.Getenv("MAIL_SENDER_EMAIL"),
		Password: os.Getenv("MAIL_SENDER_PASSWORD"),
	}
	email.StartEmailScheduler(emailService, emailCfg)

	// run app
	port := os.Getenv("PORT")
	r.Run("0.0.0.0:" + port)
}
