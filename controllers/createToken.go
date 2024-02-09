package controllers

import (
    "github.com/golang-jwt/jwt"
    "time"
    M "docker/models"
)

func createToken(user M.User) (string, error) {
    // تعیین کلید مخفی برای امضای توکن
    var secretKey = []byte("your_secret_key")

    // تعیین اطلاعات مورد نیاز برای توکن
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["user"] = user
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // مهلت انقضا 24 ساعت

    // ایجاد توکن JWT
    signedToken, err := token.SignedString(secretKey)
    if err != nil {
        return "", err
    }

    return signedToken, nil
}
