package events

import (
	"fmt"
	"log"

	"gopkg.in/gomail.v2"
)

func createUserEmail(data interface{}) {
	fmt.Println("create user email", data)

	// data2 := map[string]interface{}{
	// 	"Name":     "PR",
	// 	"Age":      30,
	// 	"Location": "New York",
	// }

	// mail.SendEmail("create_user", "create_user", data2, "pudthaiiii@gmail.com", "")

	sendEmail()
}

func sendEmail() error {
	m := gomail.NewMessage()

	// Set email sender and recipient.
	m.SetHeader("From", "your-email@example.com")
	m.SetHeader("To", "recipient@example.com")

	// Optionally add more recipients or CC/BCC
	// m.SetHeader("Cc", "cc-recipient@example.com")

	// Set the email subject.
	m.SetHeader("Subject", "Test Email from Go")

	// Set the email body.
	m.SetBody("text/plain", "This is a test email sent from Go.")

	// Attach a file (optional).
	// m.Attach("/path/to/file")

	// SMTP server configuration.
	d := gomail.NewDialer("sandbox.smtp.mailtrap.io", 2525, "a6d540ba1d45e7", "626985bd1f73c4")

	// Send the email.
	if err := d.DialAndSend(m); err != nil {
		log.Println(err)
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}
