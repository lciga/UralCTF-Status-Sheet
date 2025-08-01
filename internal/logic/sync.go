package logic

import (
	"log"

	"UralCTF-Status-Sheet/internal/gitlab"
	"UralCTF-Status-Sheet/internal/tables"
	"google.golang.org/api/sheets/v4"
)

// Запись задачи в Google Sheets
func SyncTaskEntry(srv *sheets.Service, spreadsheetID string, writeRange string, projectID string, category string, taskName string, openMR gitlab.MergeRequest) {
	// Пытаемся достать challenge.yml
	data := gitlab.GetYAML(projectID, openMR, taskName, category)
	hasYAML := true

	// Парсим YAML
	task := gitlab.Task{}
	if hasYAML {
		task = gitlab.ParseTask(data)
	} else {
		log.Printf("\033[31mФайл challenge.yml не найден для %s/%s\033[0m", category, taskName)
	}
	log.Printf("Задача %s/%s: %s", category, taskName, task.Description)

	// Определяем статус
	status := DetermineStatus(hasYAML, openMR)
	log.Printf("Статус задачи %s/%s: %s", category, taskName, status)

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
	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{row},
	}
	tables.WriteDataToSheet(srv, spreadsheetID, writeRange, valueRange)
	log.Printf("Запись задачи %s/%s в таблицу", category, taskName)
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
