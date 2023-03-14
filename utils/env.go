package utils

import (
	"fmt"
	"os"

	env "github.com/joho/godotenv"
)

func Env(key string) string {
	// load .env file
	err := env.Load(".env")
	if err != nil {
		value, ok := os.LookupEnv(key)
		if !ok {
			panic(fmt.Sprintf("can't get env variable: %s", key))
		}
		return value
	}
	return os.Getenv(key)
}

// func GetEnv(c *fiber.Ctx) error {
// 	secretKey, err := GetEnvVar("SESSION_SECRET_KEY")
// 	if err != nil {
// 		c.SendString("get env. variable error")
// 	}
// 	return c.SendString(secretKey)
// }
