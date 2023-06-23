package utils

import (
	"bytes"
	"fmt"
	"net/smtp"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug"
	"github.com/google/uuid"
)

func AddGenericVariables(variables fiber.Map) fiber.Map {
	genericVariables := fiber.Map{"BASE_URL": os.Getenv("BASE_URL")}
	merged := fiber.Map{}
	for k, v := range genericVariables {
		merged[k] = v
	}
	for k, v := range variables {
		merged[k] = v
	}

	return merged
}

func SendMail(template string, to string, subject string, variables fiber.Map) error {
	renderingEngine := pug.New("templates/emails", ".pug")
	from := os.Getenv("SMTP_FROM")
	user := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	var buf bytes.Buffer
	err := renderingEngine.Render(&buf, template, AddGenericVariables(variables))
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=\"utf-8\"\r\nDate: %s\r\nMessage-Id: %s\r\nContent-Transfer-Encoding: quoted-printable\r\nContent-Disposition: inline\r\n\r\n%s", from, to, subject, time.Now().Format(time.RFC822), uuid.New().String(), buf.String())
	auth := smtp.PlainAuth("", user, password, host)
	return smtp.SendMail(host+":"+port, auth, from, []string{to}, []byte(msg))
}
