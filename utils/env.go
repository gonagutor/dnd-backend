package utils

import (
	"log"
	"os"

	"github.com/TwiN/go-color"
	"github.com/joho/godotenv"
)

func SetupEnv() {
	err := godotenv.Load(".env")
	if err != nil && os.Getenv("enviroment") != "production" {
		log.Println(color.Ize(color.Yellow, "No .env file. Copy .env.example to .env or create a new one. You can ignore this message if running on a docker container"))
	}

	if os.Getenv("POSTGRES_PASSWORD") == "" {
		log.Fatalf(color.Ize(color.Red, "POSTGRES_PASSWORD missing from enviroment"))
	}
	if os.Getenv("POSTGRES_USER") == "" {
		log.Fatalf(color.Ize(color.Red, "POSTGRES_USER missing from enviroment"))
	}
	if os.Getenv("POSTGRES_DB") == "" {
		log.Fatalf(color.Ize(color.Red, "POSTGRES_DB missing from enviroment"))
	}
	if os.Getenv("POSTGRES_PORT") == "" {
		log.Fatalf(color.Ize(color.Red, "POSTGRES_PORT missing from enviroment"))
	}
	if os.Getenv("POSTGRES_HOST") == "" {
		log.Fatalf(color.Ize(color.Red, "POSTGRES_HOST missing from enviroment"))
	}
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatalf(color.Ize(color.Red, "JWT_SECRET missing from enviroment"))
	}

	if os.Getenv("MONGO_PASSWORD") == "" {
		log.Fatalf(color.Ize(color.Red, "MONGO_PASSWORD missing from enviroment"))
	}
	if os.Getenv("MONGO_USERNAME") == "" {
		log.Fatalf(color.Ize(color.Red, "MONGO_USERNAME missing from enviroment"))
	}
	if os.Getenv("MONGO_PORT") == "" {
		log.Fatalf(color.Ize(color.Red, "MONGO_PORT missing from enviroment"))
	}
	if os.Getenv("MONGO_HOST") == "" {
		log.Fatalf(color.Ize(color.Red, "MONGO_HOST missing from enviroment"))
	}

	if os.Getenv("BASE_URL") == "" {
		log.Fatalf(color.InYellow("BASE_URL missing from enviroment"))
	}

	if os.Getenv("SMTP_FROM") == "" {
		log.Println(color.InYellow("SMTP_FROM missing from enviroment"))
	}
	if os.Getenv("SMTP_USER") == "" {
		log.Println(color.InYellow("SMTP_USER missing from enviroment"))
	}
	if os.Getenv("SMTP_PASSWORD") == "" {
		log.Println(color.InYellow("SMTP_PASSWORD missing from enviroment"))
	}
	if os.Getenv("SMTP_HOST") == "" {
		log.Println(color.InYellow("SMTP_HOST missing from enviroment"))
	}
	if os.Getenv("SMTP_PORT") == "" {
		log.Println(color.InYellow("SMTP_PORT missing from enviroment"))
	}
}
