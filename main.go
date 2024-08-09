package main

import (
	_ "dnd/backend/docs/swagger"
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
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/pug"
)

func SetupApp() {
	viewEngine := pug.New("templates", ".pug")
	app := fiber.New(fiber.Config{
		Views:   viewEngine,
		Prefork: false,
	})

	app.Use(helmet.New())
	app.Use(logger.New(logger.ConfigDefault))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3001, http://127.0.0.1:3001" + os.Getenv("ALLOWED_ORIGINS"),
	}))

	app.Static("/static", "./static")
	app.Get("/metrics", monitor.New(monitor.Config{Title: "DND Metrics Page"}))
	app.Get("/swagger/*", swagger.HandlerDefault)

	routes.SetupVersionedRoutes(app)
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
	err = utils.PGConnection.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatal(color.InRed("Could not auto migrate db: ") + err.Error())
	}
}

// @title DND Api
// @version 1.0
// @description This is the API for https://dnd.gonagutor.com. This API handles content from the DND books, characters, campaigns and users
// @host localhost:8080
// @BasePath /
// @schemes http https

// @tag.name Auth
// @tag.description The auth system uses an access token that expires every 15 minutes and a refresh token

// @contact.name   Gonzalo Aguado Torres
// @contact.url    https://dnd.gonagutor.com/support
// @contact.email  gonagutor@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	utils.SetupEnv()
	utils.SetupPostgresConnection()
	validators.SetupValidator()
	if !fiber.IsChild() {
		AutoMigrateAll()
	}
	SetupApp()
}
