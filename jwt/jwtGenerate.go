package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)
func GenerateToken(userId uint, secretKey string) (string, error) {
    claims := jwt.MapClaims{
        "userId": userId,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }
    return tokenString, nil
}