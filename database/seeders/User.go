package seeders

import (
	D "docker/database"
	M "docker/models"
)

func UserSeeder() {
	D.DB().Create(M.User{
		Name:     "مصطفی",
		Email:    "mostafa@gmail.com",
		Password: "This is my password",
		Role:     1,
	})
}
