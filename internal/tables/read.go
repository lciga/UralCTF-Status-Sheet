package tables

import (
	"log"

	"google.golang.org/api/sheets/v4"
)

// Чтение данных из таблицы
func ReadDataFromSheet(srv *sheets.Service, spreadsheetId string, readRange string) [][]interface{} {
	read, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Ошибка чтения данных из таблицы %s: %v", spreadsheetId, err)
	}

	return read.Values
}
