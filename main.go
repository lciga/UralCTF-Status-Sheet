package main

import (
	"UralCTF-Status-Sheet/internal/gitlab"
	"fmt"
	"log"
)

func main() {
	mergeRequests, err := gitlab.GetMergeRequests("2", "opened")
	if err != nil {
		log.Fatalf("Ошибка получения merge requests: %v", err)
	}
	if len(mergeRequests) == 0 {
		log.Println("Нет открытых merge requests")
		return
	}

	rawTask, err := gitlab.GetYAML("2", mergeRequests, "too_many_redirects", "web")
	if err != nil {
		log.Fatalf("Ошибка получения YAML: %v", err)
	}

	task, err := gitlab.ParseTask(rawTask)
	if err != nil {
		log.Fatalf("Ошибка печати задачи: %v", err)
	}
	fmt.Println(task)
}
