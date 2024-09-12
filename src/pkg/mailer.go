package pkg

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"strings"
)

type emailClient struct {
	smtpServer  string
	smtpPort    string
	username    string
	password    string
	fromAddress string
	encryption  string
	TemplateDir string
}

func NewEmailClient() *emailClient {
	return &emailClient{
		smtpServer:  os.Getenv("MAIL_HOST"),
		smtpPort:    os.Getenv("MAIL_PORT"),
		username:    os.Getenv("MAIL_USERNAME"),
		password:    os.Getenv("MAIL_PASSWORD"),
		fromAddress: os.Getenv("MAIL_FROM_ADDRESS"),
		encryption:  os.Getenv("MAIL_ENCRYPTION"),
		TemplateDir: "src/mails",
	}
}

func (e *emailClient) renderTemplate(templateFile string, data map[string]interface{}) (string, error) {
	tmpl, err := template.ParseFiles(e.TemplateDir + "/" + templateFile)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (e *emailClient) SendEmail(subject, templateFile string, data map[string]interface{}, toEmail string, bccEmails ...string) error {
	body, err := e.renderTemplate(templateFile, data)
	if err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	auth := smtp.PlainAuth("", e.username, e.password, e.smtpServer)

	recipients := []string{toEmail}
	if len(bccEmails) > 0 {
		recipients = append(recipients, bccEmails...)
	}

	msg := []byte(
		fmt.Sprintf(
			"From: %s\nTo: %s\nBcc: %s\nSubject: %s\nContent-Type: text/html; charset=\"UTF-8\"\n\n%s",
			e.fromAddress, toEmail, formatBCC(bccEmails), subject, body,
		),
	)

	send := e.getSendFunction()

	err = send(auth, recipients, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func formatBCC(bccEmails []string) string {
	if len(bccEmails) == 0 {
		return ""
	}
	return fmt.Sprintf("%s", bccEmails)
}

func (e *emailClient) getSendFunction() func(auth smtp.Auth, recipients []string, msg []byte) error {
	switch strings.ToLower(e.encryption) {
	case "tls":
		return e.sendEmailTLS
	default:
		return e.sendEmailPlain
	}
}

func (e *emailClient) sendEmailPlain(auth smtp.Auth, recipients []string, msg []byte) error {
	return smtp.SendMail(
		e.smtpServer+":"+e.smtpPort,
		auth,
		e.fromAddress,
		recipients,
		msg,
	)
}

func (e *emailClient) sendEmailTLS(auth smtp.Auth, recipients []string, msg []byte) error {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         e.smtpServer,
	}

	conn, err := tls.Dial("tcp", e.smtpServer+":"+e.smtpPort, tlsConfig)
	if err != nil {
		return fmt.Errorf("failed to dial TLS connection: %w", err)
	}

	client, err := smtp.NewClient(conn, e.smtpServer)
	if err != nil {
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}

	defer client.Quit()
	err = client.Auth(auth)
	if err != nil {
		return fmt.Errorf("failed to authenticate: %w", err)
	}

	err = client.Mail(e.fromAddress)
	if err != nil {
		return fmt.Errorf("failed to set sender address: %w", err)
	}

	for _, recipient := range recipients {
		err = client.Rcpt(recipient)
		if err != nil {
			return fmt.Errorf("failed to add recipient: %w", err)
		}
	}

	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("failed to create data writer: %w", err)
	}

	_, err = w.Write(msg)
	if err != nil {
		return fmt.Errorf("failed to write message: %w", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("failed to close data writer: %w", err)
	}

	return nil
}
