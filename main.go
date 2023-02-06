package main

import (
	"log"

	"github.com/fernandovmedina/ApiRestWithJWT/database/mysql"
	"github.com/fernandovmedina/ApiRestWithJWT/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Nos conectamos a la base de datos antes de comenzar la app
	mysql.ConnectDB()

	// Creamos una app de Fiber
	app := fiber.New()

	// Mandamos a llamar todas las rutas de la app
	handlers.SetupApp(app)

	// Mandamos a ejecutar la app en una ruta en especifico
	defer func() {
		log.Fatal(app.Listen(":8080"))
	}()
}
