package email

import (
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	smtp string
	port int
	host string
	pwd  string
}

func NewEmailSender() *EmailSender {
	return &EmailSender{
		smtp: viper.GetString("Email.stmp"),
		port: viper.GetInt("Email.port"),
		host: viper.GetString("Email.host"),
		pwd:  viper.GetString("Email.pwd"),
	}
}

func (e *EmailSender) dialAndSender(message *gomail.Message) error {
	d := gomail.NewDialer(e.smtp, e.port, e.host, e.pwd)
	if err := d.DialAndSend(message); err != nil {
		return err
	}
	return nil
}

func (e *EmailSender) Send(to string, title, content string) error {
	message := gomail.NewMessage()
	message.SetHeader("From", e.host)
	message.SetHeader("To", to)
	message.SetHeader("Subject", title)
	message.SetBody("text/html", content)
	return e.dialAndSender(message)
}
