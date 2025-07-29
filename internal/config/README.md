# Пакет config
Отвечает за первичную конфигурацию проекта: работу с ,env файлами и создание сервиса для работы с Google Sheets.

## `func InitEnv()`
Инициализирует файл .env из директории проекта и загружает переменные окружения.
## `func GetEnv(key string) string`
Используется для для получения переменных окружения из файла .env. На вход прнимает имя переменной, которую необходимо извлечь и возвращает знаачение переменной.

## `func credsExtraction() *google.Credentials`
Извлекает учётные данные из JSON-файла ключа, находящегося по пути, хранящемуся в переменной окружения. Возвращает указатель на объект типа [google.Credentials](https://pkg.go.dev/golang.org/x/oauth2@v0.30.0/google#Credentials) - учётные данные для сервиса Google.
## `func ServiceCreation() *sheets.Service`
Используя извлечённые учётные данные создаёт сервис для работы с Google Sheets. Возвращает указатель на объект типа [sheets.Service](https://pkg.go.dev/google.golang.org/api@v0.243.0/sheets/v4#Service) - сервис Google Sheets.