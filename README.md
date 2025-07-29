# UralCTF Status Sheet

Данный проект реализует синхронизацию между GitLab репозиторием и Google Sheets для ведения учёта прогресса работы над тасками UralCTF.

## Реализация
Вся логика работы программы реализована в пакетах:
- [config]() - основная конфигурация программы
- [gitlab]() - работа с API GitLab
- [logic]() - основная логика работы программы
- [tables]() - работа с Google Sheets

## Запуск
Перед запуском необхрходимо склонировать репозиторий и установить зависимости:
```sh
git clone https://github.com/lciga/UralCTF-Status-Sheet.git
cd UralCTF-Status-Sheet
go mod tidy
```
Нужно создать проект и пользователя с правами Editor Google Cloud Console. Для пользователя создать ключ и скачать его в формате JSON. Также необходимо создать токен GitLab с разрешениями минимум `read_api`, `read_repository`

В корне проекта создаём файл .env следующей структуры:
```
PATH_TO_KEY="..."
SPREADSHEET_ID="..."
c="..."
GITLAB_PROJECT_ID="..."
GITLAB_URL="..."
```
где, 
- `PATH_TO_KEY` - путь к файлу Google Service Account JSON 
- `SPREADSHEET_ID` - [ID таблицы](https://developers.google.com/workspace/sheets/api/guides/concepts)
- `GITLAB_TOKEN` - [токен GitLab](https://docs.gitlab.com/api/rest/authentication/)
- `GITLAB_PROJECT_ID` - ID проекта в GitLab
- `GITLAB_URL` - адрес GitLab

Запусти программу
```sh
go run main.go
```


## TODO
- Протестировать парсинг YAML
- Протестировать синхронизацию таблицы
- Мелкие правки, касающиеся извлечения данных из .env
- Написать в `main.go` итоговыйвый вариант функции `main`