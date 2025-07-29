package main

import (
	"UralCTF-Status-Sheet/internal/gitlab"
	"fmt"
	"log"
)

func main() {
	mr, err := gitlab.GetMergeRequests("2", "all")
	if err != nil {
		log.Fatalf("Ошибка получения merge requests: %v", err)
	}
	fmt.Println(mr)
	fmt.Println("-------------------------------------------------------------------------------")

	commits, err := gitlab.GetCommit("2", "main", "tasks/web/too_many_redirects")
	if err != nil {
		log.Fatalf("Ошибка получения коммитов: %v", err)
	}
	fmt.Println(commits)
	fmt.Println("-------------------------------------------------------------------------------")

	tasks, err := gitlab.GetTasks("2", "web")
	if err != nil {
		log.Fatalf("Ошибка получения тасков: %v", err)
	}
	fmt.Println(tasks)
	fmt.Println("-------------------------------------------------------------------------------")
}
