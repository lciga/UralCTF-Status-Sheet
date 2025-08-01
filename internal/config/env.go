package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Загрузка переменных окружения из .env файла
func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("\033[31mОшибка чтения .env файла: %v\033[0m", err)
	}
	log.Println("Переменные окружения успешно загружены")
}

// Получение значения переменной окружения по ключу
func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("\033[31mПеременная окружения %v не установлена\033[0m", key)
	}
	log.Printf("Переменная окружения %v успешно получена", key)

	return value
}
