package main

import (
	"task/db"
	"task/routes"
)

func main() {
	db.Connect()              // Инициализация БД
	r := routes.SetupRouter() // Инициализация роутера
	r.Run(":8080")            // Запуск сервера
}
