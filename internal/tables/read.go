package tables

import (
	"log"

	"google.golang.org/api/sheets/v4"
)

// Чтение данных из таблицы
func ReadDataFromSheet(srv *sheets.Service, spreadsheetId string, readRange string) [][]interface{} {
	read, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Printf("\033[31mОшибка чтения данных из таблицы %s: %v\033[0m", spreadsheetId, err)
	}

	return read.Values
}
