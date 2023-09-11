package service

import (
	"bytes"
	"embed"
	"fmt"
	"net/smtp"
	"text/template"

	"github.com/msalbrain/authSphere/internals/database"
)

type MailService interface {
	SendNewUserMail(receiverName, receiverEmail string) error
	GetNewUserMessage() (string, error)
	UpdateNewUserMessage() error
}

type PlainMailConfig struct {
	host     string
	port     string
	username string
	password string
}

type MailSmtpService struct {
	smtp.Auth
	MailConfig        PlainMailConfig
	defaultFileSystem embed.FS
	q                 *database.Queries
	config            Config
}

// SendNewUserMessage implements MailService.
func (*MailSmtpService) SendNewUserMessage(receiverName string, receiverEmail string) error {
	panic("unimplemented")
}

func NewSmtpMailService(mailFs embed.FS, q *database.Queries, env Config) *MailSmtpService {
	return &MailSmtpService{smtp.PlainAuth(env.APPLICATION_NAME, env.Mail.Username, env.Mail.Password, env.Mail.Host),
		PlainMailConfig{env.Mail.Host, env.Mail.Port, env.Mail.Username, env.Mail.Password}, mailFs, q, env}
}

func (mail *MailSmtpService) GetAppName() string {
	return mail.config.APPLICATION_NAME
}

func (mail *MailSmtpService) GetMailTemplateFromDb(int) {

}

func (mail *MailSmtpService) GetConfirmLink() string {
	return ""
}

func (mail *MailSmtpService) SendNewUserMail(receiverName, receiverEmail string) error {
	subject := fmt.Sprintf("Confirm Your Signup - %s", mail.GetAppName())
	templateName := "templates/template1.html"

	templateContent, err := mail.defaultFileSystem.ReadFile(templateName)

	pureContent, err := mail.GetNewUserMessage()

	t := template.New("template")
	if string(templateContent) == "" {
		t, _ = t.Parse(pureContent)
	} else {
		t, _ = t.Parse(string(templateContent))
	}

	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))

	t.Execute(&body, struct {
		Name             string
		ApplicationName  string
		ConfirmationLink string
	}{
		Name:             receiverName,
		ApplicationName:  mail.config.APPLICATION_NAME,
		ConfirmationLink: mail.GetConfirmLink(),
	})

	// Sending email.
	err = smtp.SendMail(mail.MailConfig.host+":"+mail.MailConfig.port, mail.Auth, mail.MailConfig.username, []string{receiverEmail}, body.Bytes())

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent!")
	return nil
}

func (mail *MailSmtpService) GetNewUserMessage() (string, error) {
	return "", nil
}

func (mail *MailSmtpService) UpdateNewUserMessage() error {
	return nil
}
