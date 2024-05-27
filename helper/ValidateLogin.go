package helper

import (
	"database/sql"
	"errors"
	"net/http"
)

type User struct {
    Id       int
    Username string
    Email    string
    Password string
    Token    sql.NullString
}

func ValidateUserLogin(db *sql.DB, username, password string, w http.ResponseWriter, r *http.Request) (*User, error) {
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

    if !user.Token.Valid || user.Token.String == "" {
        http.Error(w, "Silahkan login", http.StatusUnauthorized)
        return nil, errors.New("silahkan login")
    }

    if user.Id == 0 || user.Username == "" || user.Email == "" || user.Password == "" {
        http.Error(w, "Invalid user data", http.StatusInternalServerError)
        return nil, errors.New("invalid user data")
    }

    if user.Token.Valid && user.Token.String != "" {
        http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
        return user, nil
    }

    http.Redirect(w, r, "/login", http.StatusSeeOther)
    return nil, errors.New("silahkan login")
}