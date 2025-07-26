package email

import (
	"fmt"
	"net/smtp"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SendEmail(cfg Config, to, subject, body string, isHtml bool) error {
	contentType := "text/plain"
	if isHtml {
		contentType = "text/html"
	}

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: %s; charset=\"UTF-8\"\r\n\r\n%s",
		cfg.Sender, to, subject, contentType, body)

	auth := smtp.PlainAuth("", cfg.Sender, cfg.Password, cfg.SMTPHost)
	addr := fmt.Sprintf("%s:%s", cfg.SMTPHost, cfg.SMTPPort)

	return smtp.SendMail(addr, auth, cfg.Sender, []string{to}, []byte(msg))
}

func (s *Service) SendPendingEmails(cfg Config) {
	fmt.Println("Start sending pending emails...")

	emails, err := s.repo.FindAllPendingEmail()
	if err != nil {
		fmt.Println("Error fetching queue: ", err)
		return
	}

	for _, email := range emails {
		fmt.Println("Sending email with ID: ", email.ID)
		err = s.SendEmail(cfg, email.Recipient, email.Subject, email.Body, email.IsHTML)
		if err != nil {
			err = s.repo.UpdateQueueFailed(email.ID, err.Error())
			if err != nil {
				fmt.Println("Error updating failed queue: ", err)
			}
		} else {
			err = s.repo.UpdateQueueSuccess(email.ID)
			if err != nil {
				fmt.Println("Error updating success queue: ", err)
			}
		}
	}

	fmt.Println("End sending pending emails...")
}
