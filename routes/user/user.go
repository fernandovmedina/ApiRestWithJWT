package user

import (
	"github.com/fernandovmedina/ApiRestWithJWT/database/mysql"
	"github.com/fernandovmedina/ApiRestWithJWT/models/user"
	"github.com/gofiber/fiber/v2"
)

// Funcion para obtener todos los users
func GetUsers(c *fiber.Ctx) error {
	// Creamos un arreglo que contenga todos los Users de la base de datos
	var users []user.User
	// Buscamos todos los users en la base de datos
	var result = mysql.DB.Model(&user.User{}).Find(&users)
	// Comprobamos que se haya encontrado por minimo 1 registro
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "users not found or no registers yet",
			"data":    users,
		})
	}
	// Si no se ejecuta el condicional anterior significa que si se encontraron los users en la base de datos
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    users,
	})
}

// Funcion para obtener un unico user
func GetUser(c *fiber.Ctx) error {
	// Obtenemos el id por url
	var userId = c.Params("id")
	// Creamos una variable que guarde todos los datos del user
	var user = new(user.User)
	// Guardamos el resultado en una variable
	var result = mysql.DB.Find(&user, userId)
	// Comprobamos que se haya encontrado el user
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
			"data":    nil,
		})
	}
	// Si no se ejecuta el anterior condicional, significa que si se encontro el user
	return c.Status(fiber.StatusFound).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

// Funcion para crear un user
func PostUser(c *fiber.Ctx) error {
	// Creamos una variable que guarde todos los datos de la variable
	var user = new(user.User)
	// Obtenemos todos los datos que han sido mandados por json
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "user not created, json not acceptable",
			"data":    nil,
		})
	}
	// Creamos el user y lo guardamos en la base de datos
	mysql.DB.Create(&user)
	// Si no ocurrio algun error al momento de obtener los datos por json
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

// Funcion para editar y actualizar un user
func PatchUser(c *fiber.Ctx) error {
	// Obtenemos el id por url
	var userId = c.Params("id")
	// Creamos una variable que guarde los datos de user
	var user = new(user.User)
	// Obtenemos los datos pasados por formato json
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "user not updated",
			"data":    nil,
		})
	}
	// Actualizamos el usuario y lo guardamos en la base de datos
	mysql.DB.Where("id = ?", userId).Updates(&user)
	// Si no se ejecuto el anterior condicional significa que no ocurrio ningun error al momento de ejecucion
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

// Funcion para borrar un user
func DeleteUser(c *fiber.Ctx) error {
	// Obtenemos el id pasado por url
	var userId = c.Params("id")
	// Creamos una variable que guarde todos los datos del usuario
	var user = new(user.User)
	// Guardamos el resultado de borrar el user en una variable
	var result = mysql.DB.Delete(&user, userId)
	// Comprobamos que si se haya borrado el user
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "user not deleted",
			"data":    nil,
		})
	}
	// Si no se ejecuto el anterior condicional significa que no ocurrio ningun error al momento de ejecucion
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user deleted",
		"data":    user,
	})
}
