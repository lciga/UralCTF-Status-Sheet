# UralCTF Status Sheet

Данный проект реализует синхронизацию между GitLab репозиторием и Google Sheets для ведения учёта прогресса работы над тасками UralCTF.

## Реализация
Вся логика работы программы реализована в пакетах:
- [config](https://github.com/lciga/UralCTF-Status-Sheet/blob/main/internal/config/README.md) - основная конфигурация программы
- [gitlab](https://github.com/lciga/UralCTF-Status-Sheet/blob/main/internal/gitlab/README.md) - работа с API GitLab
- [logic](https://github.com/lciga/UralCTF-Status-Sheet/blob/main/internal/logic/README.md) - основная логика работы программы
- [tables](https://github.com/lciga/UralCTF-Status-Sheet/blob/main/internal/tables/README.md) - работа с Google Sheets

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
GITLAB_TOKEN="..."
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

## Тест парсинга YAML
```go
	mergeRequests, err := gitlab.GetMergeRequests("2", "opened")
	if err != nil {
		log.Fatalf("Ошибка получения merge requests: %v", err)
	}
	if len(mergeRequests) == 0 {
		log.Println("Нет открытых merge requests")
		return
	}

	rawTask, err := gitlab.GetYAML("2", mergeRequests, "too_many_redirects", "web")
	if err != nil {
		log.Fatalf("Ошибка получения YAML: %v", err)
	}

	task, err := gitlab.ParseTask(rawTask)
	if err != nil {
		log.Fatalf("Ошибка печати задачи: %v", err)
	}
	fmt.Println(task)
```

## TODO
- Протестировать парсинг YAML
- Протестировать синхронизацию таблицы
- Мелкие правки, касающиеся извлечения данных из .env
- Написать в `main.go` итоговыйвый вариант функции `main`
- Переработать логику работы программы: выдавать статус "в прогрессе" при создании ветки (написать функцию получения веток???), как вытаскивать challenge.md???