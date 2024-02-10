package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/golang-jwt/jwt"
)

var (
	Store    *session.Store
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)

// TokenDetails struct to hold access and refresh token details
type TokenDetails struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AtExpires    int64  `json:"accessTokenExpiresAt"`
	RtExpires    int64  `json:"refreshTokenExpiresAt"`
}

var JWTSecretKey = []byte(Env("ACCESS_SECRET_KEY"))
var JWTRefreshSecretKey = []byte(Env("REFRESH_SECRET_KEY"))

var accessTokenLifetime = EnvToInt("ACCESS_TOKEN_LIFETIME")
var refreshTokenLifetime = EnvToInt("REFRESH_TOKEN_LIFETIME")

// ! generate a 45 long character session id
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
func Session(c *fiber.Ctx) *session.Session {
	sess, err := Store.Get(c)
	if err != nil {
		panic(err)
	}
	return sess
}

// CreateTokens generates new access and refresh tokens for a given user
func CreateToken(userID uint) (*TokenDetails, error) {
	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * time.Duration(accessTokenLifetime)).Unix() // Access token expires after 15 minutes
	td.RtExpires = time.Now().Add(time.Hour * time.Duration(refreshTokenLifetime)).Unix()  // Refresh token expires after 7 days

	var err error
	// Access token
	atClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userID,
		"exp":        td.AtExpires,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString(JWTSecretKey)
	if err != nil {
		return nil, err
	}

	// Refresh token
	rtClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     td.RtExpires,
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString(JWTRefreshSecretKey)
	if err != nil {
		return nil, err
	}

	return td, nil
}
