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