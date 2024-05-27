package context

import (
	"database/sql"
	"errors"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

var secretKey = []byte(os.Getenv("TOKEN_SECRET"))

func GetUserIdFromContext(db *sql.DB, r *http.Request) (int, error) {
	userID := r.Context().Value("user_id")
	if userID == nil {
		return 0, errors.New("user_id not found in request context")
	}

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return 0, errors.New("authorization header not found")
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("invalid user_id in token claims")
	}
	var userId = int(userIDFloat)


	query := `SELECT user_id FROM users WHERE token = $1`
	var userIDFromDB int
	err = db.QueryRow(query, tokenString).Scan(&userIDFromDB)
	if err != nil {
		return 0, err
	}

	if userIDFromDB != userID {
		return 0, errors.New("user_id in token does not match user_id in database")
	}

	return userId, nil
}