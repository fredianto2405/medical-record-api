package email

type Config struct {
	SMTPHost string
	SMTPPort string
	Sender   string
	Password string
}

type DTO struct {
	ID        int    `db:"id"`
	Recipient string `db:"recipient"`
	Subject   string `db:"subject"`
	Body      string `db:"body"`
	IsHTML    bool   `db:"is_html"`
}
