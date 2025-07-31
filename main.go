package main

import (
	"UralCTF-Status-Sheet/internal/config"
	"UralCTF-Status-Sheet/internal/gitlab"
	"UralCTF-Status-Sheet/internal/logic"

	"log"
)

func main() {
	config.InitEnv()

	category := []string{"web", "pwn", "crypto", "stego", "forensic", "reverse", "ppc"}

	spreadsheetID := config.GetEnv("SPREADSHEET_ID")
	projectID := config.GetEnv("GITLAB_PROJECT_ID")
	srv := config.ServiceCreation()
	// logic.SyncTaskEntry(srv, spreadsheetID, projectID, "web", "too_many_redirects", gitlab.GetMergeRequests(projectID, "all"))

	for _, c := range category {
		task := gitlab.GetTasks(projectID, c)
		for i := range len(task) {
			logic.SyncTaskEntry(srv, spreadsheetID, projectID, c, task[i].Name, gitlab.GetMergeRequests(projectID, "all"))
			log.Printf("Синхронизации %s категория %s", task[i].Name, c)
		}
	}
}
