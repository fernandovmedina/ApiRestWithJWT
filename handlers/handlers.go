package handlers

import (
	authRoute "github.com/fernandovmedina/ApiRestWithJWT/controllers/auth"
	apiRoute "github.com/fernandovmedina/ApiRestWithJWT/routes/api"
	userRoute "github.com/fernandovmedina/ApiRestWithJWT/routes/user"
	fiber "github.com/gofiber/fiber/v2"
)

func SetupApp(app *fiber.App) {
	////////////////////////////////////////////////////////

	/*
		API
	*/
	// Creamos el grupo principal => /api
	api := app.Group("/api")
	// Fijamos una ruta en => /api == /api/
	api.Get("/", apiRoute.Hello)

	////////////////////////////////////////////////////////

	/*
		AUTH
	*/
	// Creamos un grupo en => /api/auth
	auth := api.Group("/auth")
	// Fijamos con un metodo POST a la ruta de ingreso de la api
	auth.Post("/login", authRoute.Login)
	// Fijamos con un metodo POST a la ruta de registro de la api
	auth.Post("/register", authRoute.Register)
	// Fijamos con un metodo POST a la ruta de cerrar sesion de la api
	auth.Post("/logout", authRoute.Logout)
	// Fijamos con un metodo GET a la ruta de user
	auth.Get("/user", authRoute.User)

	////////////////////////////////////////////////////////

	/*
		USER
	*/
	// Creamos un grupo en => /api/user
	user := api.Group("/user")
	// Fijamos una ruta con Metodo GET, para obtener todos los usuarios
	user.Get("/", userRoute.GetUsers)
	// Fijamos una ruta con Metodo GET, para obtener un unico usuario
	user.Get("/:id", userRoute.GetUser)
	// Fijamos una ruta con Metodo POST, para crear un usuario
	user.Post("/", userRoute.PostUser)
	// Fijamos una ruta con Metodo PATCH, para editar un usuario
	user.Patch("/:id", userRoute.PatchUser)
	// Fijamos una ruta con Metodo DELETE, para borrar un usuario
	user.Delete("/:id", userRoute.DeleteUser)

	////////////////////////////////////////////////////////
}
