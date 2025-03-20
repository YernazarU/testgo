package config

import (
    "database/sql"
    _ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
    connStr := "postgres://user:password@localhost:5432/users_db?sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }
    
    err = db.Ping()
    if err != nil {
        return nil, err
    }
    
    return db, nil
}