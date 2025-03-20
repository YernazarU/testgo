package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "user-api/config"
    "user-api/handlers"
)

func main() {
    db, err := config.InitDB()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    r := mux.NewRouter()
    
    r.HandleFunc("/users", handlers.CreateUser(db)).Methods("POST")
    r.HandleFunc("/users/{id}", handlers.GetUser(db)).Methods("GET")
    r.HandleFunc("/users/{id}", handlers.UpdateUser(db)).Methods("PUT")

    log.Println("Server starting on :8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}