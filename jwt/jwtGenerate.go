package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string, email string, id uint, secretKey string) (string, error) {
    claims := jwt.MapClaims{
        "username": username,
        "email": email,
        "id": id,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }
    return tokenString, nil
}