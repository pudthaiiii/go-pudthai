package mailer

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"go-ibooking/internal/config"
	"go-ibooking/internal/infrastructure/logger"
	"html/template"
	"net/smtp"
	"strings"
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

func (e *Mailer) renderTemplate(templateFile string, data map[string]interface{}) (string, error) {
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

func (e *Mailer) SendEmail(subject, templateFile string, data map[string]interface{}, toEmail string, bccEmails ...string) error {
	// body, err := e.renderTemplate(templateFile, data)
	// if err != nil {
	// 	logger.Log.Err(err).Msg("failed to render template")
	// 	return fmt.Errorf("failed to render template: %w", err)
	// }

	auth := smtp.PlainAuth("", e.username, e.password, e.smtpServer)

	recipients := []string{toEmail}
	if len(bccEmails) > 0 {
		recipients = append(recipients, bccEmails...)
	}

	msg := []byte(
		fmt.Sprintf(
			"From: %s\nTo: %s\nBcc: %s\nSubject: %s\nContent-Type: text/html; charset=\"UTF-8\"\n\n%s",
			e.fromAddress, toEmail, formatBCC(bccEmails), subject, "<h1>hello</h1>",
		),
	)

	fmt.Println(e.smtpServer, e.smtpPort, e.username, e.password, e.fromAddress, e.encryption, e.TemplateDir)
	sendErr := smtp.SendMail(
		e.smtpServer+":"+e.smtpPort,
		auth,
		e.fromAddress,
		recipients,
		msg,
	)

	if sendErr != nil {
		logger.Log.Err(sendErr).Msg("failed to send email")
		return fmt.Errorf("failed to send email: %w", sendErr)
	}

	fmt.Println(auth, recipients, msg)

	// send := e.getSendFunction()

	// err = send(auth, recipients, msg)
	// if err != nil {
	// 	logger.Log.Err(err).Msg("failed to send email")
	// 	return fmt.Errorf("failed to send email: %w", err)
	// }

	logger.Write.Info().Msg("Email sent successfully")

	return nil
}

func formatBCC(bccEmails []string) string {
	if len(bccEmails) == 0 {
		return ""
	}

	return fmt.Sprintf("%s", bccEmails)
}

func (e *Mailer) getSendFunction() func(auth smtp.Auth, recipients []string, msg []byte) error {
	fmt.Println(e.encryption)
	switch strings.ToLower(e.encryption) {
	case "tls":
		return e.sendEmailTLS
	default:
		return e.sendEmailPlain
	}
}

func (e *Mailer) sendEmailPlain(auth smtp.Auth, recipients []string, msg []byte) error {
	return smtp.SendMail(
		e.smtpServer+":"+e.smtpPort,
		auth,
		e.fromAddress,
		recipients,
		msg,
	)
}

func (e *Mailer) sendEmailTLS(auth smtp.Auth, recipients []string, msg []byte) error {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         e.smtpServer,
	}

	conn, err := tls.Dial("tcp", e.smtpServer+":"+e.smtpPort, tlsConfig)
	if err != nil {
		logger.Log.Err(err).Msg("failed to dial TLS connection")
		return fmt.Errorf("failed to dial TLS connection: %w", err)
	}

	client, err := smtp.NewClient(conn, e.smtpServer)
	if err != nil {
		logger.Log.Err(err).Msg("failed to create SMTP client")
		return fmt.Errorf("failed to create SMTP client: %w", err)
	}

	defer client.Quit()
	err = client.Auth(auth)
	if err != nil {
		logger.Log.Err(err).Msg("failed to authenticate")
		return fmt.Errorf("failed to authenticate: %w", err)
	}

	err = client.Mail(e.fromAddress)
	if err != nil {
		logger.Log.Err(err).Msg("failed to set sender address")
		return fmt.Errorf("failed to set sender address: %w", err)
	}

	for _, recipient := range recipients {
		err = client.Rcpt(recipient)
		if err != nil {
			logger.Log.Err(err).Msg("failed to add recipient")
			return fmt.Errorf("failed to add recipient: %w", err)
		}
	}

	w, err := client.Data()
	if err != nil {
		logger.Log.Err(err).Msg("failed to create data writer")
		return fmt.Errorf("failed to create data writer: %w", err)
	}

	_, err = w.Write(msg)
	if err != nil {
		logger.Log.Err(err).Msg("failed to write message")
		return fmt.Errorf("failed to write message: %w", err)
	}

	err = w.Close()
	if err != nil {
		logger.Log.Err(err).Msg("failed to close data writer")
		return fmt.Errorf("failed to close data writer: %w", err)
	}

	return nil
}
