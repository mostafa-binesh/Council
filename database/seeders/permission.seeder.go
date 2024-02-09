package seeders

import (
	D "docker/database"
	M "docker/models"
)

func PermissionSeeder() {
	permissions := []string{
		"view-admin",
		"view-user",
		"edit",
		"create",
		"delete",
	}
	for i := 0; i < len(permissions); i++ {
		D.DB().Create(&M.Permission{
			Name: permissions[i],
		})
	}
}