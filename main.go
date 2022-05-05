package main

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	keyauth "github.com/iwpnd/fiber-key-auth"
	"log"
	"sign-builder/core"
	"sign-builder/handlers"
)

func main() {
	core.Init()
	//Make sure fonts are available when script is loaded
	_ = rice.MustFindBox("fonts")
	_ = rice.MustFindBox("templates")

	app := fiber.New()

	//Logging
	app.Use(logger.New())

	// Called through Hasura
	api := app.Group("/api", keyauth.New())
	api.Get("/getshield", handlers.HandleShieldQuery)
	api.Post("/saveshield", handlers.HandleShieldPostQuery)

	app.Get("/health", handlers.HealthHandler)

	log.Fatal(app.Listen(":3000"))
}
