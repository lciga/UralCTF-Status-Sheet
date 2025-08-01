package config

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var ctx = context.Background()

// Экстракция кредов из JSON файла ключа
func credsExtraction() *google.Credentials {
	// Извлечение переменных из .env
	keysPath := GetEnv("PATH_TO_KEY")

	// Чтенеие JSON файла с кредами
	jsonCreds, err := os.ReadFile(keysPath)
	if err != nil {
		log.Printf("\033[31mОшибка чтения файла ключа: %v\033[0m", err)
	}
	log.Println("Файл ключа успешно прочитан")

	// Извлечение кредов из JSON и получение объекта *google.Credentials
	creds, err := google.CredentialsFromJSON(ctx, jsonCreds, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Printf("\033[31mОшибка извлечения учётных данных: %v\033[0m", err)
	}
	log.Println("Учётные данные успешно извлечены")

	return creds
}

// Создание сервиса Google Sheets API
func ServiceCreation() *sheets.Service {
	credsExtraction()
	srv, err := sheets.NewService(ctx, option.WithCredentials(credsExtraction()))
	if err != nil {
		log.Printf("\033[31mОшибка создания сервиса Google Sheets API: %v\033[0m", err)
	}
	log.Println("Сервис Google Sheets API успешно создан")

	return srv
}
