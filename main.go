package main

import (
	C "docker/controllers"
	D "docker/database"
	R "docker/routes"
	U "docker/utils"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	D.ConnectToDB() // initialize database
	C.Initilize()   // initialize controllers value
	R.RouterInit()
	// ! session
	U.Store = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 5,
		KeyGenerator: func() string {
			// secretKey, err := C.GetEnvVar("SESSION_SECRET_KEY")
			secretKey := U.Env("SESSION_SECRET_KEY")
			var sessionID string
			sessionID, err := U.GenerateSessionID(secretKey)
			if err != nil {
				panic(err)
			}
			return sessionID
		},
	})

}
