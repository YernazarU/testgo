package config

import (
    "database/sql"
    "log"
    "time"
    _ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
    connStr := "postgres://user:password@db:5432/users_db?sslmode=disable" // Изменено localhost на db
    var db *sql.DB
    var err error

    for i := 0; i < 10; i++ {
        db, err = sql.Open("postgres", connStr)
        if err != nil {
            log.Printf("Failed to open DB: %v", err)
            time.Sleep(2 * time.Second)
            continue
        }

        err = db.Ping()
        if err == nil {
            log.Println("Successfully connected to database")
            return db, nil
        }
        
        log.Printf("Failed to ping DB: %v, retrying in 2s...", err)
        time.Sleep(2 * time.Second)
    }

    return nil, err
}