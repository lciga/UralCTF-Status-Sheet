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
		log.Fatalf("Ошибка чтения .env файла: %v", err)
	}
	log.Println("Переменные окружения успешно загружены")
}

// Получение значения переменной окружения по ключу
func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Переменная окружения %v не установлена", key)
	}
	log.Printf("Переменная окружения %v успешно получена", key)

	return value
}
