package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "user-api/models"
)

func CreateUser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var user models.User
        json.NewDecoder(r.Body).Decode(&user)
        
        err := db.QueryRow(
            "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
            user.Name, user.Email,
        ).Scan(&user.ID)
        
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        
        json.NewEncoder(w).Encode(user)
    }
}

func GetUser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]
        
        var user models.User
        err := db.QueryRow(
            "SELECT id, name, email FROM users WHERE id = $1",
            id,
        ).Scan(&user.ID, &user.Name, &user.Email)
        
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        
        json.NewEncoder(w).Encode(user)
    }
}

func UpdateUser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        id := vars["id"]
        
        var user models.User
        json.NewDecoder(r.Body).Decode(&user)
        
        _, err := db.Exec(
            "UPDATE users SET name = $1, email = $2 WHERE id = $3",
            user.Name, user.Email, id,
        )
        
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        
        json.NewEncoder(w).Encode(user)
    }
}