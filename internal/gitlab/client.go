package gitlab

import (
	"log"
	"net/http"

	"UralCTF-Status-Sheet/internal/config"
)

var Client *http.Client
var BaseURL string
var Token string

// Инициализация клиента GitLab
func InitClient() {
	Token = config.GetEnv("GITLAB_TOKEN")
	BaseURL = config.GetEnv("GITLAB_URL")
	Client = &http.Client{}
}

// Отправка запроса к GitLab API
func SendRequest(path string) *http.Response {
	InitClient()

	// Создание запроса для получения тасков
	req, err := http.NewRequest("GET", BaseURL+path, nil)
	if err != nil {
		log.Fatalf("Ошибка создания запроса: %v", err)
	}
	req.Header.Set("Private-Token", Token)
	log.Printf("Отправка запроса: %s", req.URL)

	// Отправка запроса
	resp, err := Client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка отправки ответа: %v", err)
	}
	log.Printf("Получен ответ: %s", resp.Status)

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка получения данных: %v", resp.Status)
	}

	return resp
}
