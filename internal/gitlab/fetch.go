package gitlab

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type MergeRequest struct {
	ID          int       `json:"id"`
	Iid         int       `json:"iid"`
	ProjectID   int       `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	State       string    `json:"state"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	MergedBy    struct {
		ID          int    `json:"id"`
		Username    string `json:"username"`
		PublicEmail any    `json:"public_email"`
		Name        string `json:"name"`
		State       string `json:"state"`
		Locked      bool   `json:"locked"`
		AvatarURL   string `json:"avatar_url"`
		WebURL      string `json:"web_url"`
	} `json:"merged_by"`
	MergeUser struct {
		ID          int    `json:"id"`
		Username    string `json:"username"`
		PublicEmail any    `json:"public_email"`
		Name        string `json:"name"`
		State       string `json:"state"`
		Locked      bool   `json:"locked"`
		AvatarURL   string `json:"avatar_url"`
		WebURL      string `json:"web_url"`
	} `json:"merge_user"`
	MergedAt       time.Time `json:"merged_at"`
	ClosedBy       any       `json:"closed_by"`
	ClosedAt       any       `json:"closed_at"`
	TargetBranch   string    `json:"target_branch"`
	SourceBranch   string    `json:"source_branch"`
	UserNotesCount int       `json:"user_notes_count"`
	Upvotes        int       `json:"upvotes"`
	Downvotes      int       `json:"downvotes"`
	Author         struct {
		ID          int    `json:"id"`
		Username    string `json:"username"`
		PublicEmail any    `json:"public_email"`
		Name        string `json:"name"`
		State       string `json:"state"`
		Locked      bool   `json:"locked"`
		AvatarURL   string `json:"avatar_url"`
		WebURL      string `json:"web_url"`
	} `json:"author"`
	Assignees []struct {
		ID          int    `json:"id"`
		Username    string `json:"username"`
		PublicEmail any    `json:"public_email"`
		Name        string `json:"name"`
		State       string `json:"state"`
		Locked      bool   `json:"locked"`
		AvatarURL   string `json:"avatar_url"`
		WebURL      string `json:"web_url"`
	} `json:"assignees"`
	Assignee struct {
		ID          int    `json:"id"`
		Username    string `json:"username"`
		PublicEmail any    `json:"public_email"`
		Name        string `json:"name"`
		State       string `json:"state"`
		Locked      bool   `json:"locked"`
		AvatarURL   string `json:"avatar_url"`
		WebURL      string `json:"web_url"`
	} `json:"assignee"`
	Reviewers []struct {
		ID          int    `json:"id"`
		Username    string `json:"username"`
		PublicEmail any    `json:"public_email"`
		Name        string `json:"name"`
		State       string `json:"state"`
		Locked      bool   `json:"locked"`
		AvatarURL   string `json:"avatar_url"`
		WebURL      string `json:"web_url"`
	} `json:"reviewers"`
	SourceProjectID           int       `json:"source_project_id"`
	TargetProjectID           int       `json:"target_project_id"`
	Labels                    []any     `json:"labels"`
	Draft                     bool      `json:"draft"`
	Imported                  bool      `json:"imported"`
	ImportedFrom              string    `json:"imported_from"`
	WorkInProgress            bool      `json:"work_in_progress"`
	Milestone                 any       `json:"milestone"`
	MergeWhenPipelineSucceeds bool      `json:"merge_when_pipeline_succeeds"`
	MergeStatus               string    `json:"merge_status"`
	DetailedMergeStatus       string    `json:"detailed_merge_status"`
	MergeAfter                any       `json:"merge_after"`
	Sha                       string    `json:"sha"`
	MergeCommitSha            any       `json:"merge_commit_sha"`
	SquashCommitSha           string    `json:"squash_commit_sha"`
	DiscussionLocked          any       `json:"discussion_locked"`
	ShouldRemoveSourceBranch  bool      `json:"should_remove_source_branch"`
	ForceRemoveSourceBranch   bool      `json:"force_remove_source_branch"`
	PreparedAt                time.Time `json:"prepared_at"`
	Reference                 string    `json:"reference"`
	References                struct {
		Short    string `json:"short"`
		Relative string `json:"relative"`
		Full     string `json:"full"`
	} `json:"references"`
	WebURL    string `json:"web_url"`
	TimeStats struct {
		TimeEstimate        int `json:"time_estimate"`
		TotalTimeSpent      int `json:"total_time_spent"`
		HumanTimeEstimate   any `json:"human_time_estimate"`
		HumanTotalTimeSpent any `json:"human_total_time_spent"`
	} `json:"time_stats"`
	Squash               bool `json:"squash"`
	SquashOnMerge        bool `json:"squash_on_merge"`
	TaskCompletionStatus struct {
		Count          int `json:"count"`
		CompletedCount int `json:"completed_count"`
	} `json:"task_completion_status"`
	HasConflicts                bool `json:"has_conflicts"`
	BlockingDiscussionsResolved bool `json:"blocking_discussions_resolved"`
	ApprovalsBeforeMerge        any  `json:"approvals_before_merge"`
}

type Commit struct {
}

func GetCommit(projectID string) ([]Commit, error) {
	return nil, nil
}

func GetMergeRequests(projectID string) ([]MergeRequest, error) {
	InitClient()

	// Создание запроса для получения merge requests
	req, err := NewRequest(http.MethodGet, "api/v4/projects/"+projectID+"/merge_requests")
	if err != nil {
		log.Fatalf("Ошибка создания запроса: %v", err)
	}
	log.Printf("Отправка запроса: %s", req.URL)

	// Отправка запроса
	resp, err := Client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка отправки ответа: %v", err)
	}
	defer resp.Body.Close()
	log.Printf("Получен ответ: %s", resp.Status)

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Ошибка получения данных: %v", resp.Status)
	}
	log.Println("Декодирование ответа")

	// Декодирование ответа в срез структур MergeRequest
	var mergeRequests []MergeRequest
	if err := json.NewDecoder(resp.Body).Decode(&mergeRequests); err != nil {
		log.Fatalf("Ошибка декодирования ответа: %v", err)
	}
	log.Printf("Получено %d merge requests", len(mergeRequests))

	return mergeRequests, nil
}

func GetYAML() {

}
