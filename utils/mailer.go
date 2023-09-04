package utils

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"math"
	"math/big"
	"net/smtp"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/pug"
)

func addGenericVariables(variables fiber.Map) fiber.Map {
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

func generateMessageID() string {
	t := time.Now().UnixNano()
	pid := os.Getpid()
	rint, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	h, err := os.Hostname()
	if err != nil {
		h = "localhost"
	}
	msgid := fmt.Sprintf("<%d.%d.%d@%s>", t, pid, rint, h)
	return msgid
}

func SendMail(template string, to string, subject string, variables fiber.Map) error {
	renderingEngine := pug.New("templates/emails", ".pug")
	from := os.Getenv("SMTP_FROM")
	user := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")
	useSSL := os.Getenv("SMTP_USE_SSL")

	var buf bytes.Buffer
	err := renderingEngine.Render(&buf, template, addGenericVariables(variables))
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=\"utf-8\"\r\nDate: %s\r\nMessage-Id: %s\r\nContent-Transfer-Encoding: quoted-printable\r\nContent-Disposition: inline\r\n\r\n%s", from, to, subject, time.Now().Format(time.RFC1123Z), generateMessageID(), buf.String())

	auth := smtp.PlainAuth("", user, password, host)
	if useSSL == "true" {
		tlsconfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         host,
		}
		conn, err := tls.Dial("tcp", host+":"+port, tlsconfig)
		if err != nil {
			return err
		}

		client, err := smtp.NewClient(conn, host)
		if err != nil {
			return err
		}

		err = client.Auth(auth)
		if err != nil {
			return err
		}

		err = client.Mail(from)
		if err != nil {
			return err
		}

		err = client.Rcpt(to)
		if err != nil {
			return err
		}

		writer, err := client.Data()
		if err != nil {
			return err
		}

		_, err = writer.Write([]byte(msg))
		if err != nil {
			return err
		}

		err = writer.Close()
		if err != nil {
			return err
		}

		return client.Quit()
	}

	return smtp.SendMail(host+":"+port, auth, from, []string{to}, []byte(msg))
}
