package seeders

import (
	D "docker/database"
	M "docker/models"
)

func PermissionSeeder() {
	permissions := []string{
		"ViweLaw",
		"AddLaw",
		"EditLaw",
		"DeleteLaw",
		"ViweAllComments",
		"AccOrRefComments",
		"ViweUser",
		"AddUser",
		"EditUser",
		"DeleteUser",
		"AccOrRefUser",
	}
	RandomId := []string{
		"8df73d8c-aea8-40e9-be50-bd5c339b2cbe",
		"a5a6836f-6d01-42c9-b573-5902247e0bf6",
		"992647f6-72cc-4b0b-9fae-059e17f90717",
		"2ab0dc45-23fc-4309-a873-1e75ee862131",
		"6d3d7f53-d86e-4573-ad2d-cac88d512fec",
		"e517b102-5fe8-4d2c-9839-e1ccd68685a3",
		"bbcab700-ab9d-43fb-9920-8ee83b2e2b03",
		"5f493046-8c84-4bfa-98f2-3fe35fc2d6d3",
		"be6c6dfd-0fda-4ce5-ba6a-ee69ed52aa00",
		"a775b3ec-c110-4561-9ca6-b9644479436c",
		"2d41daf7-1d61-476d-af25-764fc62a662f",
	}
	for i := 0; i < len(permissions); i++ {
		D.DB().Create(&M.Permission{
			Name:     permissions[i],
			RandomID: RandomId[i],
		})
	}
}
