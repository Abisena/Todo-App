package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo-app/auth/service"
	"todo-app/database"
)


var register string = "Ayoo Register\n"
func Register(users *service.User, w http.ResponseWriter, r *http.Request) {
	var user service.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.ServiceUserRegis(&user, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

    fmt.Printf("%v", register)
}


func Login(users *service.User, w http.ResponseWriter, r *http.Request) {
    var user service.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    newDB, err := database.ConnectDatabase()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer newDB.Close()

    loggedUser, err := service.ServiceUserLogin(newDB, user.Username, user.Password, w, r)
    if err != nil {
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(loggedUser)
}