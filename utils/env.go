package utils

import (
	env "github.com/joho/godotenv"
	"os"
)

func Env(key string) (string, error) {
	// load .env file
	err := env.Load(".env")
	if err != nil {
		return "", err
	}
	return os.Getenv(key), nil
}

// func GetEnv(c *fiber.Ctx) error {
// 	secretKey, err := GetEnvVar("SESSION_SECRET_KEY")
// 	if err != nil {
// 		c.SendString("get env. variable error")
// 	}
// 	return c.SendString(secretKey)
// }
