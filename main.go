package main

import (
	"UralCTF-Status-Sheet/internal/config"
	"UralCTF-Status-Sheet/internal/tables"
	"google.golang.org/api/sheets/v4"
)

func main() {
	config.InitEnv()
	srv := config.ServiceCreation()

	data := &sheets.ValueRange{
		Values: [][]interface{}{
			{"В процессе"},
		},
	}

	tables.WriteDataToSheet(srv, config.GetEnv("SPREADSHEET_ID"), config.GetEnv("SHEET_NAME")+"!A5", data)
}
