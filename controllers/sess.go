package controllers

// ! session controller
import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// ! generate a 45 long character generate id
// ! generated by chatGPT
func GenerateSessionID(secret string) (string, error) {
	// Generate 32 bytes of random data
	idBytes := make([]byte, 32)
	if _, err := rand.Read(idBytes); err != nil {
		return "", err
	}

	// Append the secret code to the bytes
	idBytes = append(idBytes, []byte(secret)...)

	// Hash the bytes using SHA-256
	hashBytes := sha256.Sum256(idBytes)

	// Encode the hash using base64
	id := base64.StdEncoding.EncodeToString(hashBytes[:])

	return id, nil
}
// ! get session from fiber context
func GetSess(c *fiber.Ctx) *session.Session {
	sess, err := Store.Get(c)
	if err != nil {
		panic(err)
	}
	return sess
}
