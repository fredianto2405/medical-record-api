package email

import (
	"fmt"
	"medical-record-api/pkg/logger"
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
	log := logger.Log
	log.Info("========== Start sending pending emails ==========")

	emails, err := s.repo.FindAllPendingEmail()
	if err != nil {
		log.WithError(err).Error("Error fetching queue: ")
		return
	}

	if len(emails) == 0 {
		log.Info("No pending emails found")
		log.Info("========== End sending pending emails ==========")
		return
	}

	log.Infof("Found %d pending emails", len(emails))

	for _, email := range emails {
		log.WithFields(map[string]interface{}{
			"queue_id": email.ID,
			"to":       email.Recipient,
			"subject":  email.Subject,
		}).Info("Sending email")

		err = s.SendEmail(cfg, email.Recipient, email.Subject, email.Body, email.IsHTML)
		if err != nil {
			log.WithFields(map[string]interface{}{
				"queue_id": email.ID,
				"to":       email.Recipient,
				"error":    err.Error(),
			}).Error("Failed to send email")

			errUpdate := s.repo.UpdateQueueFailed(email.ID, err.Error())
			if errUpdate != nil {
				log.Error("Error updating failed queue: ", errUpdate.Error())
			}
		} else {
			log.WithFields(map[string]interface{}{
				"queue_id": email.ID,
				"to":       email.Recipient,
			}).Info("Email sent successfully")

			errUpdate := s.repo.UpdateQueueSuccess(email.ID)
			if errUpdate != nil {
				log.Error("Error updating success queue: ", errUpdate.Error())
			}
		}
	}

	log.Info("========== End sending pending emails ==========")
}
