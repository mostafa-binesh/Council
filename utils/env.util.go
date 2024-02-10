package utils

import (
	"fmt"
	"os"
	"strconv"

	env "github.com/joho/godotenv"
)

// returns envoirment variable if exist in the .env file
// otherwise, try to get the env. variable from host
// if there was no env. variable, panics
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
func EnvToBool(key string) bool {
	value := Env(key) // Use your Env function to get the value as a string
	if value == "true" {
		return true
	}
	if value == "false" {
		return false
	}
	panic(fmt.Sprintf("env variable not a bool: %s", key))
}
func EnvToInt(key string) int {
	value := Env(key) // Use your Env function to get the value as a string
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(fmt.Sprintf("can't convert env variable to int: %s, error: %v", key, err))
	}
	return intValue
}
