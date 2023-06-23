package utils

import (
	"log"

	"github.com/TwiN/go-color"
	"github.com/joho/godotenv"
)

func SetupEnv() {
	godotenv.Load(".env")
	env, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalf(color.Ize(color.Red, "No .env file. Copy .env.example to .env or create a new one"))
	} else {
		if env["POSTGRES_PASSWORD"] == "" {
			log.Fatalf(color.Ize(color.Red, "POSTGRE_PASSWORD missing on .env file"))
		}
		if env["POSTGRES_USER"] == "" {
			log.Fatalf(color.Ize(color.Red, "POSTGRES_USER missing on .env file"))
		}
		if env["POSTGRES_DB"] == "" {
			log.Fatalf(color.Ize(color.Red, "POSTGRES_DB missing on .env file"))
		}
		if env["POSTGRES_PORT"] == "" {
			log.Fatalf(color.Ize(color.Red, "POSTGRES_PORT missing on .env file"))
		}
		if env["POSTGRES_HOST"] == "" {
			log.Fatalf(color.Ize(color.Red, "POSTGRES_HOST missing on .env file"))
		}
		if env["JWT_SECRET"] == "" {
			log.Fatalf(color.Ize(color.Red, "JWT_SECRET missing on .env file"))
		}
	}
}
