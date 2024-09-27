package mailer

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"go-ibooking/internal/config"
	"go-ibooking/internal/infrastructure/logger"
	"html/template"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
)

type Mailer struct {
	smtpServer  string
	smtpPort    string
	username    string
	password    string
	fromAddress string
	encryption  string
	TemplateDir string
}

func NewMailer(cfg *config.Config) *Mailer {
	return &Mailer{
		smtpServer:  cfg.Get("MailServer")["SmtpServer"].(string),
		smtpPort:    cfg.Get("MailServer")["SmtpPort"].(string),
		username:    cfg.Get("MailServer")["Username"].(string),
		password:    cfg.Get("MailServer")["Password"].(string),
		fromAddress: cfg.Get("MailServer")["FromAddress"].(string),
		encryption:  cfg.Get("MailServer")["Encryption"].(string),
		TemplateDir: "internal/mails",
	}
}

func (e *Mailer) renderTemplate(templateFile string, data any) (string, error) {
	templateFile = templateFile + ".html"

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

func (e *Mailer) Send(subject, templateFile string, data any, toEmail string, bccEmails ...string) error {
	body, err := e.renderTemplate(templateFile, data)
	if err != nil {
		logger.Log.Err(err).Msg("failed to render template")
		return fmt.Errorf("failed to render template: %w", err)
	}

	m := gomail.NewMessage()

	m.SetHeader("From", e.fromAddress)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	if len(bccEmails) > 0 {
		m.SetHeader("Bcc", bccEmails...)
	}

	port, err := strconv.Atoi(e.smtpPort)
	if err != nil {
		logger.Log.Err(err).Msg("invalid SMTP port")
		return fmt.Errorf("invalid SMTP port: %w", err)
	}

	d := gomail.NewDialer(
		e.smtpServer,
		port,
		e.username,
		e.password,
	)

	if strings.ToLower(e.encryption) == "tls" {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	if err := d.DialAndSend(m); err != nil {
		logger.Log.Err(err).Msg("failed to send email")
		return err
	}

	logger.Write.Info().Msg("Email sent successfully")
	return nil
}
