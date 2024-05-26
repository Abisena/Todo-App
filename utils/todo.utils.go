package utils

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"todo-app/database"

	"github.com/gorilla/mux"
)

var TesksUtils = "Ini Utils Data Crud"

type Todo struct {
	Id string
	Judul string
	Tanggal string
	Acara string
}

func UtilsTodoCrud(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	switch r.Method {
	case http.MethodPost:
		var user Todo
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = createDataTodo(db, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	case http.MethodGet:
		todos, err := getAllDataTodo(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(todos)
	case http.MethodPut:
		var user Todo
		err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = updateDataTodo(db, &user, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		id, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = deleteDataTodo(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func createDataTodo(db *sql.DB, user *Todo) error {
	query := `INSERT INTO todo (judul, tanggal, acara) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, user.Judul, user.Tanggal, user.Acara)
	return err
}


func getAllDataTodo(db *sql.DB) ([]Todo, error) {
	rows, err := db.Query("SELECT * FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Judul, &todo.Tanggal, &todo.Acara)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func updateDataTodo(db *sql.DB, user *Todo, id int) error {
	_, err := db.Exec("UPDATE todo SET judul = $1, tanggal = $2, acara = $3 WHERE id = $4", user.Judul, user.Tanggal, user.Acara, id)
	return err
}

func deleteDataTodo(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM todo WHERE id = $1", id)
	return err
}