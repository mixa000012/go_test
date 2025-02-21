package service

import (
	"awesomeProject/config"
	"awesomeProject/db"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type UserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResponse struct {
	ID int `json:"id"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	dbConn := db.InitPostgres(config.Cfg.PGConnStr)
	defer dbConn.Close()

	if err := createUsersTable(dbConn); err != nil {
		http.Error(w, "Ошибка при создании таблицы", http.StatusInternalServerError)
		return
	}

	var userID int
	err := dbConn.QueryRow("INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id", req.Name, req.Age).Scan(&userID)
	if err != nil {
		http.Error(w, "Ошибка при вставке в базу", http.StatusInternalServerError)
		return
	}

	resp := UserResponse{ID: userID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func createUsersTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INT NOT NULL
	);`
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("Ошибка создания таблицы: %v", err)
		return err
	}
	return nil
}
