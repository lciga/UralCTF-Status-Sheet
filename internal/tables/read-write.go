package tables

import (
	"log"

	"google.golang.org/api/sheets/v4"
)

// Запись данных в таблицу
func writeDataToSheet(srv *sheets.Service, spreadsheetId string, data *sheets.ValueRange, writeRange string) {
	_, err := srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, data).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Ошибка добавления данных в таблицу: %s", err)
	}
	log.Println("Данные успешно добавлены в таблицу")
}

// Чтение данных из таблицы
func readDataFromSheet(srv *sheets.Service, spreadsheetId string, readRange string) [][]interface{} {
	read, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatal(err)
	}

	return read.Values
}
