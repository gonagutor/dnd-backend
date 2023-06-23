package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/TwiN/go-color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var PGConnection *gorm.DB

func GetConnectionPG() (*gorm.DB, error) {
	var logMode logger.LogLevel = logger.Silent

	user, _ := os.LookupEnv("POSTGRES_USER")
	password, _ := os.LookupEnv("POSTGRES_PASSWORD")
	database, _ := os.LookupEnv("POSTGRES_DB")
	host, _ := os.LookupEnv("POSTGRES_HOST")
	port, _ := os.LookupEnv("POSTGRES_PORT")
	logSql, _ := os.LookupEnv("LOG_SQL")

	switch logSql {
	case "info":
		logMode = logger.Info
	case "warn":
		logMode = logger.Warn
	case "error":
		logMode = logger.Error
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, database, port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logMode)})
}

func SetupPostgresConnection() {
	conn, pgerr := GetConnectionPG()
	if pgerr != nil {
		log.Fatal(color.InRed("Could not connect to db: ") + pgerr.Error())
	}
	_, dberr := conn.DB()
	if dberr != nil {
		log.Fatal(color.InRed("Could not connect to db: ") + dberr.Error())
	}
	PGConnection = conn
}
