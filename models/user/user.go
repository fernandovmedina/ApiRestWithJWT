package user

import "gorm.io/gorm"

// Creamos una estructura para los Users
type User struct {
	gorm.Model
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}
