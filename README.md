# User API

REST API для управления пользователями, написанное на Go с использованием PostgreSQL.

## Требования

- Docker
- Docker Compose
- Go 1.21 (для локальной разработки, если не используете Docker)

## Запуск

1. Склонируйте репозиторий:
git clone <github-url>
cd user-api


2. Убедитесь, что у вас есть файлы `go.mod` и `go.sum`. Если их нет, выполните: 
go mod init user-api
go get github.com/gorilla/mux
go get github.com/lib/pq


3. Запустите проект с помощью Docker Compose:   
    docker-compose up --build

- Это соберет образ приложения и запустит контейнеры (приложение на порту 8080, PostgreSQL на порту 5432).
- Таблица `users` будет создана автоматически при первом запуске базы данных.

4. API будет доступно на `http://localhost:8080`.

## Эндпоинты

- `POST /users` - создать пользователя (пример: `{"name": "Arthur", "email": "goodman@morgan.com"}`)
- `GET /users/{id}` - получить пользователя
- `PUT /users/{id}` - обновить пользователя (пример: `{"name": "Kobe Bryant", "email": "kobe.bean@nba.com"}`)

## Тестирование

Используйте curl или Postman:
curl -X POST -H "Content-Type: application/json" -d '{"name": "Arthur", "email": "goodman@morgan.com"}' http://localhost:8080/users
curl http://localhost:8080/users/1