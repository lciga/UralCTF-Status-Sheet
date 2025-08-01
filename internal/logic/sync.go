package logic

import (
	"log"

	"UralCTF-Status-Sheet/internal/gitlab"
	"UralCTF-Status-Sheet/internal/tables"
	"google.golang.org/api/sheets/v4"
)

// Запись задачи в Google Sheets
func SyncTaskEntry(srv *sheets.Service, spreadsheetID string, projectID string, category string, taskName string, openMR gitlab.MergeRequest) {
	// Пытаемся достать challenge.yml
	data, err := gitlab.GetYAML(projectID, openMR, taskName, category)
	hasYAML := err == nil

	// Парсим YAML
	task := gitlab.Task{}
	if hasYAML {
		task, err = gitlab.ParseTask(data)
		if err != nil {
			log.Printf("Ошибка парсинга YAML для %s: %v", taskName, err)
			hasYAML = false
		}
	} else {
		log.Printf("Файл challenge.yml не найден для %s/%s", category, taskName)
	}

	// Определяем статус
	status := DetermineStatus(hasYAML, openMR)

	// Формируем строку
	row := []interface{}{
		task.Name,                // A: Название
		task.Description,         // B: Описание для борды
		task.Category,            // C: Категория
		getDifficulty(task.Tags), // D: Сложность
		task.Author,              // E: Кто делает
		string(status),           // F: Статус
	}

	// Пишем в таблицу
	writeRange := "Tasks!A2"
	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{row},
	}
	tables.WriteDataToSheet(srv, spreadsheetID, writeRange, valueRange)
}

// Достаём сложность из тегов
func getDifficulty(tags []string) string {
	for _, tag := range tags {
		switch tag {
		case "easy", "medium", "hard":
			return tag
		}
	}
	return ""
}
