package tables

import (
	"google.golang.org/api/sheets/v4"
	"log"
)

// Запись данных в таблицу
func WriteDataToSheet(srv *sheets.Service, spreadsheetId string, writeRange string, data *sheets.ValueRange) {
	_, err := srv.Spreadsheets.Values.Update(spreadsheetId, writeRange, data).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Ошибка добавления данных в таблицу %s: %v", spreadsheetId, err)
	}
	log.Printf("Данные успешно добавлены в таблицу %s", spreadsheetId)
}
