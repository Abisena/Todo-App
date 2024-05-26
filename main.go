package main

import (
	"fmt"
	"net/http"
	"todo-app/auth"
	"todo-app/database"
	"todo-app/utils"

	"github.com/gorilla/mux"
)

func main() {
	db, err := database.ConnectDatabase()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	auth.Register()

	r := mux.NewRouter()

	r.HandleFunc("/todos", utils.UtilsTodoCrud).Methods(http.MethodPost, http.MethodGet)
	r.HandleFunc("/todos/{id}", utils.UtilsTodoCrud).Methods(http.MethodPut, http.MethodDelete)

	fmt.Println(http.ListenAndServe(":8000", r))
}
