package config

import (
    "database/sql"
    "log"
    "time"
    _ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
    connStr := "postgres://user:password@localhost:5432/users_db?sslmode=disable"
    var db *sql.DB
    var err error

    // Повторные попытки подключения
    for i := 0; i < 10; i++ { // Пробуем 10 раз
        db, err = sql.Open("postgres", connStr)
        if err != nil {
            log.Printf("Failed to open DB: %v", err)
            time.Sleep(2 * time.Second) // Ждем 2 секунды перед повторной попыткой
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

    return nil, err // Если после 10 попыток не удалось, возвращаем ошибку
}