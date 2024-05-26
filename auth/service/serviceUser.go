package service

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"todo-app/database"
	"todo-app/jwt"
)


type User struct {
	Id uint
	Username string
	Email string
	Password string
	Token string
}

func ServiceUserRegis(user *User, w http.ResponseWriter, r *http.Request) error {
    newDB, err := database.ConnectDatabase()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return err
    }
    defer newDB.Close()

    query := `INSERT INTO todo (username, email, password) VALUES ($1, $2, $3)`
    _, err = newDB.Exec(query, user.Username, user.Email, user.Password)
    return err
}


func ServiceUserLogin(db *sql.DB, username, password string, w http.ResponseWriter, r *http.Request) (*User, error) {
	query := `SELECT id, username, email, password FROM todo WHERE username = $1 AND password = $2`
	row := db.QueryRow(query, username, password)

	user := &User{}
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return nil, err
	}

	secretKey := os.Getenv("Token_SECRET")
	if secretKey == "" {
		fmt.Println("Secret key tidak ditemukan di variabel lingkungan")
		return user, err
	}

	token, err := jwt.GenerateToken(user.Username, user.Email, user.Id, secretKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	user.Token = token

	return user, nil
}