package main

import (
	"UralCTF-Status-Sheet/internal/gitlab"
	"fmt"
	"log"
)

func main() {
	mr, err := gitlab.GetMergeRequests("2")
	if err != nil {
		log.Fatalf("Ошибка получения merge requests: %v", err)
	}
	fmt.Println(mr)
}
