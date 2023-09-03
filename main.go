package main

import (
	"dnd/backend/handlers"
	"dnd/backend/models"
	"dnd/backend/routes"
	"dnd/backend/utils"
	"dnd/backend/utils/validators"
	"log"
	"os"

	"github.com/TwiN/go-color"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/template/pug"
)

func SetupApp() {
	viewEngine := pug.New("templates", ".pug")
	app := fiber.New(fiber.Config{
		Views:   viewEngine,
		Prefork: true,
	})

	app.Use(helmet.New())
	app.Use(logger.New(logger.ConfigDefault))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3001, http://127.0.0.1:3001",
	}))

	app.Static("/static", "./static")
	app.Get("/metrics", monitor.New(monitor.Config{Title: "Flat Searcher Metrics Page"}))

	routes.SetupVersionedRoutes(app)
	routes.SetupAuthRoutes(app)
	app.Get("/", handlers.Status)
	if os.Getenv("PORT") != "" {
		app.Listen(":" + os.Getenv("PORT"))
	} else {
		app.Listen(":3000")
	}
}

func AutoMigrateAll() {
	err := utils.PGConnection.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(color.InRed("Could not auto migrate db: ") + err.Error())
	}
	err = utils.PGConnection.AutoMigrate(&models.Character{})
	if err != nil {
		log.Fatal(color.InRed("Could not auto migrate db: ") + err.Error())
	}
}

func main() {
	utils.SetupEnv()
	utils.SetupPostgresConnection()
	validators.SetupValidator()
	if !fiber.IsChild() {
		AutoMigrateAll()
	}
	SetupApp()
}
