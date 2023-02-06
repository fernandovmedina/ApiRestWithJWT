package auth

import (
	"strconv"
	"time"

	"github.com/fernandovmedina/ApiRestWithJWT/database/mysql"
	"github.com/fernandovmedina/ApiRestWithJWT/models/user"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Variable que guarde la llave secreta
const SecretKey string = "fernandovmedina"

func User(c *fiber.Ctx) error {
	// Creamos una variable cookie para obtener la cookie
	var cookie = c.Cookies("jwt")
	//
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	var claims = token.Claims.(*jwt.StandardClaims)

	var user user.User

	mysql.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "succes",
		"data":    user,
	})
}

func Register(c *fiber.Ctx) error {
	// Creamos un mapa que guarde los datos del token
	var datos map[string]string
	// Convertimos los datos en formato json y los guardamos en la variable datos
	if err := c.BodyParser(&datos); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "invalid json",
		})
	}
	// Generamos el password
	password, err := bcrypt.GenerateFromPassword([]byte(datos["password"]), bcrypt.DefaultCost)
	// Verificamos que no haya ningun error al momento de generar el password
	if err != nil {
		return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"message": "there was a mistake when generating the password",
		})
	}
	// Creamos una variable que guarde los datos del user
	var user = user.User{
		Name:     datos["name"],
		LastName: datos["lastname"],
		Email:    datos["email"],
		Password: string(password),
	}
	// Creamos el usuario en la base de datos
	mysql.DB.Create(&user)
	// Retornamos la informacion en formato json
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user registered succesfully",
		"data":    user,
	})
}

func Login(c *fiber.Ctx) error {
	// Creamos un mapa que guarde todos los datos del usuario a encontrar
	var datos map[string]string
	// Verificamos que no haya ningun error al momento de obtener los datos por json
	if err := c.BodyParser(&datos); err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"message": "invalid json",
		})
	}
	// Creamos una variable de tipo user, que guarde los datos de la misma
	var user user.User
	// Buscamos el usuario en la base de datos
	mysql.DB.Where("email = ?", datos["email"]).First(&user)
	// Verificamos que el usuario existe
	if user.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}
	// Verificamos que la contraseña del usuario y la de la base de datos sea la misma
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(datos["password"])); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "incorrect password",
		})
	}
	// Generamos los claims
	var claims = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	// Asignamos la llava secreta a nuestros claims
	token, err := claims.SignedString([]byte(SecretKey))
	// Verificamos que no haya ningun error al momento de asignar la llave secreta a nuestros claims
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "could not login",
		})
	}
	// Generamos una cookie
	var cookie = fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	// Asignamos la cookie
	c.Cookie(&cookie)
	// Retornamos el siguiente mensaje y la cookie en formato json
	return c.JSON(fiber.Map{
		"message": "succesfully login",
		"cookie":  cookie,
	})
}

func Logout(c *fiber.Ctx) error {
	// Creamos una nueva cookie la cual expira instantaneamente
	var cookie = fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour * 1),
		HTTPOnly: true,
	}
	// Asignamos la cookie
	c.Cookie(&cookie)
	// Retornamos el siguiente mensaje en formato json
	return c.JSON(fiber.Map{
		"message": "success",
		"cookie":  cookie,
	})
}
