package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const defPort = "7070"

func main() {

	// Инициализируем ENV
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Установка порта по умолчанию
	port := os.Getenv("TODO_PORT")

	if port != "" {
		port = defPort
	}

	// Установка директории с фронтенд файлами
	fs := http.FileServer(http.Dir("./web"))

	// Обработка запросов к корневой директории
	http.Handle("/", fs)

	// Запуск сервера
	log.Printf("Сервер запущен на http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
