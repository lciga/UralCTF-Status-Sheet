package tables

import (
	"log"

	"google.golang.org/api/sheets/v4"
)

// Запись данных в таблицу
func WriteDataToSheet(srv *sheets.Service, spreadsheetId string, writeRange string, data *sheets.ValueRange) {
	_, err := srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, data).ValueInputOption("RAW").Do()
	if err != nil {
		log.Printf("\033[31mОшибка добавления данных в таблицу %s: %v\033[0m", spreadsheetId, err)
	}
	log.Printf("Данные успешно добавлены в таблицу %s", spreadsheetId)
}
