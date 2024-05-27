package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-app/auth"
	"todo-app/database"
	"todo-app/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	loadErr := godotenv.Load()
	if loadErr != nil {
    	log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()

	// r.Use(middleware.LoginMiddleware)

	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		auth.Register(nil, w, r)
	}).Methods(http.MethodPost)
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		auth.Login(nil, w, r)
	}).Methods(http.MethodPost)
	r.HandleFunc("/todos", utils.UtilsTodoCrud).Methods(http.MethodPost, http.MethodGet)
	r.HandleFunc("/todos/{id}", utils.UtilsTodoCrud).Methods(http.MethodPut, http.MethodDelete)

	fmt.Println(http.ListenAndServe(":8000", r))
}
