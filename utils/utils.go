package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"html/template"

	"gopkg.in/gomail.v2"
)

// EmailConfig holds SMTP credentials
type EmailConfig struct {
	From     string
	Host     string
	Port     int
	Username string
	Password string
}

// Load SMTP configuration from environment variables
func LoadEmailConfig() EmailConfig {
	return EmailConfig{
		// From:     os.Getenv("EMAIL_FROM"),
		// Host:     os.Getenv("SMTP_HOST"),
		// Port:     587, // Default SMTP port
		// Username: os.Getenv("SMTP_USER"),
		// Password: os.Getenv("SMTP_PASSWORD"),
		From:     "MS_ZXbl5m@trial-0r83ql37qzxgzw1j.mlsender.net",
		Host:     "smtp.mailersend.net",
		Port:     587,
		Username: "MS_ZXbl5m@trial-0r83ql37qzxgzw1j.mlsender.net",
		Password: "mssp.ItLhYfM.v69oxl5ywkx4785k.N1eGbz3",
	}
}

// SendEmail sends an email with dynamic HTML content
func SendEmail(to string, subject string, templateFile string, data interface{}) error {
	config := LoadEmailConfig()

	// Parse the email template
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("failed to parse email template: %w", err)
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute email template: %w", err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", config.From)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body.String())
	d := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("Email sent successfully to", to)
	return nil
}

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	// Encode to base64 to ensure it's a readable string
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}
