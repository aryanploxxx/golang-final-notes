package main

import (
	"fmt"
	"log"
	"os"
	"url-shortner/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)

}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

/*
	POST http://localhost:3000/api/v1

	{
	"url":"https://www.youtube.com/watch?v=3ExDEeSnyvE&list=PL5dTjWUk_cPbXTq9j-3Vaq08rjH1D5cTn&index=8"
	}

	{
	"url": "https://www.youtube.com/watch?v=3ExDEeSnyvE&list=PL5dTjWUk_cPbXTq9j-3Vaq08rjH1D5cTn&index=8",
	"short": "localhost:3000/de06ba",
	"expiry": 24,
	"rate_limit": 9,
	"rate_limit_reset": 30
	}
*/
