# Пакет logic
Описывает основную логику работы программы: определение статус таска и синхронизация таблицы

## `type TaskStatus string`
Тип двнных, содержащий статус таска готовности таска.
### Возможные статусы: 
- В процессе - создана директория таска с файлом challenge.yml. Информация о таске добавляется в таблицу.
- Ожидает апрува - отправлен Merge Request. Происходит проверка таска.
- Отправлен на доработку - Merge Request отклонён.
- Готов - слияние ветки таска в main.
## `func DetermineStatus(hasYAML bool, mr gitlab.MergeRequest) TaskStatus`
Предназначена для определения статуса готовности таска на основе данных о директории таска и Merge Request. Принимает на вход логическое значение - наличие YAML-файла в директории и экземпляр структуры [gitlab.MergeRequest](https://github.com/lciga/UralCTF-Status-Sheet/blob/main/internal/gitlab/README.md). Возвращает экземпляр типа TaskStatus.

## `func getDifficulty(tags []string) string`
Предназначена для парсинга сложности таска на основе тегов из файла challenge.yml. Принимает на вход массив строк. Возвращает строку.
## `func SyncTaskEntry(srv *sheets.Service, spreadsheetID string, projectID string, category string, taskName string, openMR gitlab.MergeRequest)`
Предназначена для синхронизации информации о таске в таблице. На основе данных, полученных из challenge.yml и выводов функций getDifficulty и DetermineStatus создаётся интерфейс:
```go
	row := []interface{}{
		task.Name,                // A: Название
		task.Description,         // B: Описание для борды
		task.Category,            // C: Категория
		getDifficulty(task.Tags), // D: Сложность
		task.Author,              // E: Кто делает
		string(status),           // F: Статус
	}
```
Далее данные записываются в соответствующую таблицу и в соответствующие поля. На вход податся указатель на объект типа [sheets.Service](https://pkg.go.dev/google.golang.org/api@v0.243.0/sheets/v4#Service) - сервис Google Sheets, [ID таблицы](https://developers.google.com/workspace/sheets/api/guides/concepts), ID проекта, категория и название таска, экземпляр структуры [gitlab.MergeRequest](https://github.com/lciga/UralCTF-Status-Sheet/blob/main/internal/gitlab/README.md), содержащей открытые Merge Requests.