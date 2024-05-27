package middleware

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte(os.Getenv("TOKEN_KEY"))

type contextKey string

var userIDKey contextKey = "user_id"


func AuthMiddleware(next http.Handler, db *sql.DB) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenString := r.Header.Get("Authorization")
        if tokenString == "" {
            http.Error(w, "authorization header not found", http.StatusUnauthorized)
            return
        }
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return secretKey, nil
        })
        if err != nil {
            http.Error(w, err.Error(), http.StatusUnauthorized)
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            http.Error(w, "invalid token claims", http.StatusUnauthorized)
            return
        }
        userIDFloat, ok := claims["user_id"].(float64)
        if !ok {
            http.Error(w, "invalid user_id in token claims", http.StatusUnauthorized)
            return
        }
        userID := int(userIDFloat)



        ctx := context.WithValue(r.Context(), userIDKey, userID)
        r = r.WithContext(ctx)

        next.ServeHTTP(w, r)
    })
}