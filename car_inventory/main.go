package main

import (
	"car_inventory/config"
	"car_inventory/handlers"

	"bitbucket.org/smartbot2017/go-common/logger"
	"github.com/gofiber/fiber"
)

func main() {
	config.ConnectDB()

	app := fiber.New()

	app.Use(logger.New())

	app.Post("/cars", handlers.AddCarHandler)
	/* net/http
	mux := http.NewServeMux()

	mux.HandleFunc("/cars", handlers.AddCarHandler)
	mux.HandleFunc("/cars/", handlers.CarHandler)

	wrappermux := middlewares.Logger(mux)
	wrappermux = middlewares.Security(mux)

	http.ListenAndServe(":8080", wrappermux)
	*/

}
