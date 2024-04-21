package seeders

import (
	D "docker/database"
	M "docker/models"
	"math/rand"
	"strconv"
	"time"
	"golang.org/x/crypto/bcrypt"
)

func AdminSeeder() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("Shora@Shora!1403!"), bcrypt.DefaultCost)
	if err != nil {
		
	}
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator with current time
	D.DB().Create(&M.User{
		Name:         "محمدمهدی کاظمی",
		PhoneNumber:  "09121318520",
		Password:     string(hashedPassword), // = password
		RoleID:         1,
		NationalCode: strconv.Itoa(rand.Intn(9000000000) + 1000000000), // Generate 10-digit number
		PersonalCode: "appAdminshora",                                     // Generate 10-digit number
		Verified:     true,
	})
}
