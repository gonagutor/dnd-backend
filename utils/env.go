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
		log.Println(color.Ize(color.Yellow, "No .env file. Copy .env.example to .env or create a new one. You can ignore this message if running on a docker container"))
	}

	if env["POSTGRES_PASSWORD"] == "" {
		log.Fatalf(color.Ize(color.Red, "POSTGRES_PASSWORD missing from enviroment"))
	}
	if env["POSTGRES_USER"] == "" {
		log.Fatalf(color.Ize(color.Red, "POSTGRES_USER missing from enviroment"))
	}
	if env["POSTGRES_DB"] == "" {
		log.Fatalf(color.Ize(color.Red, "POSTGRES_DB missing from enviroment"))
	}
	if env["POSTGRES_PORT"] == "" {
		log.Fatalf(color.Ize(color.Red, "POSTGRES_PORT missing from enviroment"))
	}
	if env["POSTGRES_HOST"] == "" {
		log.Fatalf(color.Ize(color.Red, "POSTGRES_HOST missing from enviroment"))
	}
	if env["JWT_SECRET"] == "" {
		log.Fatalf(color.Ize(color.Red, "JWT_SECRET missing from enviroment"))
	}

}
