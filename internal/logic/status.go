package logic

import (
	"UralCTF-Status-Sheet/internal/gitlab"
)

type TaskStatus string

const (
	StatusInProgress    TaskStatus = "В процессе"
	StatusWaitingReview TaskStatus = "Ожидает апрува"
	StatusReturned      TaskStatus = "Отправлен на доработку"
	StatusReady         TaskStatus = "Готов"
)

func DetermineStatus(hasYAML bool, mr gitlab.MergeRequest) TaskStatus {
	if mr == nil {
		if hasYAML {
			return StatusInProgress
		}
		return ""
	}

	// Берём только первый MR по ветке
	m := (mr)[0]

	if m.TargetBranch == "main" {
		switch m.State {
		case "opened":
			return StatusWaitingReview
		case "closed":
			if m.MergedAt.IsZero() {
				return StatusReturned
			}
			return StatusReady
		case "merged":
			return StatusReady
		}
	}

	// fallback: если YAML появился, но MR не в main или неизвестный статус
	if hasYAML {
		return StatusInProgress
	}

	return ""
}
