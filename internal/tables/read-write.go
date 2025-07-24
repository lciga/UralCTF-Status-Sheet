package tables

import (
	"log"

	"google.golang.org/api/sheets/v4"
)

// Запись данных в таблицу
func WriteDataToSheet(srv *sheets.Service, spreadsheetId string, writeRange string, data *sheets.ValueRange) {
	_, err := srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, data).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Ошибка добавления данных в таблицу %s: %s", spreadsheetId, err)
	}
	log.Printf("Данные успешно добавлены в таблицу %s", spreadsheetId)
}

// Чтение данных из таблицы
func ReadDataFromSheet(srv *sheets.Service, spreadsheetId string, readRange string) [][]interface{} {
	read, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Ошибка чтения данных из таблицы %s: %s", spreadsheetId, err)
	}

	return read.Values
}
