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
    Id       uint
    Username string
    Email    string
    Password string
    Token    sql.NullString
}

func ServiceUserRegis(user *User, w http.ResponseWriter, r *http.Request) error {
    newDB, err := database.ConnectDatabase()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return err
    }
    defer newDB.Close()

    query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
    _, err = newDB.Exec(query, user.Username, user.Email, user.Password)
    return err
}


func ServiceUserLogin(db *sql.DB, username, password string, w http.ResponseWriter, r *http.Request) (*User, error) {
    query := `SELECT user_id, username, email, password, token FROM users WHERE username = $1 AND password = $2`
    row := db.QueryRow(query, username, password)

    user := &User{}
    err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Password, &user.Token)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return nil, err
    }

    secretKey := os.Getenv("TOKEN_SECRET")
    if secretKey == "" {
        fmt.Println("Secret key tidak ditemukan di variabel lingkungan")
        return user, err
    }


    if !user.Token.Valid || user.Token.String == "" {
        tokenStr, err := jwt.GenerateToken(user.Id, secretKey)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return nil, err
        }
        user.Token = sql.NullString{String: tokenStr, Valid: true}

        updateQuery := `UPDATE users SET token = $1 WHERE user_id = $2`
        _, err = db.Exec(updateQuery, user.Token, user.Id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return nil, err
        }
    }

    return user, nil
}