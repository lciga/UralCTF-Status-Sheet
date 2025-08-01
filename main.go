package main

import (
	"UralCTF-Status-Sheet/internal/config"
	"UralCTF-Status-Sheet/internal/gitlab"
	"fmt"

	"UralCTF-Status-Sheet/internal/logic"
)

func main() {
	config.InitEnv()

	category := []string{"web", "pwn", "crypto", "stego", "forensic", "reverse", "ppc"}

	spreadsheetID := config.GetEnv("SPREADSHEET_ID")

	projectID := config.GetEnv("GITLAB_PROJECT_ID")
	srv := config.ServiceCreation()

	rowIndex := 2 // Начинаем со второй строки (первая — заголовки)

	for _, c := range category {
		tasks := gitlab.GetTasks(projectID, c)
		for _, t := range tasks {
			writeRange := fmt.Sprintf("Tasks!A%d:F%d", rowIndex, rowIndex)
			logic.SyncTaskEntry(srv, spreadsheetID, writeRange, projectID, c, t.Name, gitlab.GetMergeRequests(projectID, "all"))
			rowIndex++
		}
	}
}
