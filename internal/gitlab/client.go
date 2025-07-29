package gitlab

import (
	"net/http"

	"UralCTF-Status-Sheet/internal/config"
)

var Client *http.Client
var BaseURL string
var Token string

// Инициализация клиента GitLab
func InitClient() {
	config.InitEnv()
	Token = config.GetEnv("GITLAB_TOKEN")
	BaseURL = config.GetEnv("GITLAB_URL")
	Client = &http.Client{}
}

// Создание нового запроса к GitLab API
func NewRequest(method, path string) (*http.Request, error) {
	req, err := http.NewRequest(method, BaseURL+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Private-Token", Token)
	return req, nil
}
