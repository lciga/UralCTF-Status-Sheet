# UralCTF Status Sheet
TODO: Написать нормальный ридми

## Чтение данных из таблицы
```go
	spreadsheetId := os.Getenv("SPREADSHEET_ID")
	readRange := "Лист1!A:A"
	read, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(read.Values)
```
Здесь читаются данные из столбца A

## Запись данных в таблицу
```go
	data := &sheets.ValueRange{
		Values: [][]interface{}{
			{"val1", "val2", "val3"},
		},
	}

	_, err = srv.Spreadsheets.Values.Append(spreadsheetId, "Лист1!A1", data).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Ошибка добавления данных в таблицу: %s", err)
	}
```
Здесь в таблицу, начиная с ячейки A1 заполняем значениями из `data`

## Пример структуры данных, загружаемых в таблицу
```go
	data := &sheets.ValueRange{
		Values: [][]interface{}{
			{"В процессе"},
		},
	}
```
## Отправка запроса
```go
	gitlab.InitClient()

	req, err := gitlab.NewRequest(http.MethodGet, "api/v4/projects")
	if err != nil {
		log.Fatalf("Ошибка создания запроса: %v", err)
	}
	resp, err := gitlab.Client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка отправки ответа: %v", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp)
```

## Типы MR для ручки `api/v4/projects/:id/merge_request?state=:state`
- state=opened — только открытые
- state=merged — только смерженные
- state=closed — закрытые, но не смерженные
- state=all — всё подряд (default)