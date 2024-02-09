package seeders

import (
	D "docker/database"
	M "docker/models"
)

func RoleSeeder() {
	roles := []string{
		"admin",
		"super-user",
		"normally-user",
	}
	for i := 0; i < len(roles); i++ {
		D.DB().Create(&M.Role{
			Name: roles[i],
		})

	}
}
