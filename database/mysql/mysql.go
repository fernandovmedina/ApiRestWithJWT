package mysql

import (
	"log"
	"os"

	"github.com/fernandovmedina/ApiRestWithJWT/models/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Creamos una variable que guarde la conexion de la base de datos
var DB *gorm.DB

// Creamos una funcion para conectarnos a la base de datos
func ConnectDB() {
	// Creamos una variable que guarde los errores al momento de ejecutar la siguiente funcion
	// var err error
	// Cargamos el archivo .env
	var err = godotenv.Load()
	// Verificamos que no ocurra ningun error al momento de cargar el archivo .env
	if err != nil {
		panic("Ocurrio algun error al momento de cargar el archivo .env")
	}
	// Obtenemos todos los datos del archivo .env, sobre la base de datos
	var (
		databaseName     = os.Getenv("DATABASE_NAME")
		databaseUser     = os.Getenv("DATABASE_USER")
		databasePassword = os.Getenv("DATABASE_PASSWORD")
		databasePort     = os.Getenv("DATABASE_PORT")
		databaseHost     = os.Getenv("DATABASE_HOST")
	)
	// Creamos una variable qur guarde el dsn
	var dsn string = databaseUser + ":" + databasePassword + "@tcp(" + databaseHost + ":" + databasePort + ")/" + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// Abrimos la base de datos
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Printf("%s \n", "Ocurrion algun problema al momento de conectarse a la base de datos")
	} else {
		// Mostramos un mensaje por pantalla si la conexion fue exitosa
		log.Printf("%s \n", "La conexion a la base de datos fue exitosa")
		// Migramos la base de datos automaticamente
		DB.AutoMigrate(&user.User{})
	}
}
