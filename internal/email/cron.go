package email

import "github.com/robfig/cron/v3"

func StartEmailScheduler(service *Service, cfg Config) {
	c := cron.New()
	c.AddFunc("@every 1m", func() {
		service.SendPendingEmails(cfg)
	})
	c.Start()
}
