package main

import (
	"go_final_project/tests"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// Установка порта по умолчанию
	port := strconv.Itoa(tests.Port)
	// Проверка переменной окружения TODO_PORT
	if envPort, ok := os.LookupEnv("TODO_PORT"); ok {
		port = envPort
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
