package utils

import (
	"log"
	"os"

	"github.com/TwiN/go-color"
	"github.com/joho/godotenv"
)

func SetupEnv() {
	err := godotenv.Load(".env")
	if err != nil {
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

}
