package main

import (
	"car_inventory/config"
	"car_inventory/handlers"
	"car_inventory/middlewares"

	"github.com/gofiber/fiber/v2" // ✅ Correct v2 version
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config.ConnectDB()

	app := fiber.New()

	app.Use(logger.New())
	app.Use(middlewares.SecurityHeaders)
	app.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin":   "12345",
			"manager": "quality",
			"john":    "doe",
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User is not authorized",
			})
		},
	}))

	app.Post("/cars", handlers.AddCarHandler)
	app.Get("/cars/:id", handlers.GetCarHandler)
	app.Delete("/cars/:id", handlers.DeleteCarHandler)

	/* net/http
	mux := http.NewServeMux()

	mux.HandleFunc("/cars", handlers.AddCarHandler)
	mux.HandleFunc("/cars/", handlers.CarHandler)

	wrappermux := middlewares.Logger(mux)
	wrappermux = middlewares.Security(mux)

	http.ListenAndServe(":8080", wrappermux)
	*/

	app.Listen(":8080")

}
