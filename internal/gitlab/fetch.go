package gitlab

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type MergeRequest []struct {
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

type Commit []struct {
	ID             string    `json:"id"`
	ShortID        string    `json:"short_id"`
	CreatedAt      time.Time `json:"created_at"`
	ParentIds      []string  `json:"parent_ids"`
	Title          string    `json:"title"`
	Message        string    `json:"message"`
	AuthorName     string    `json:"author_name"`
	AuthorEmail    string    `json:"author_email"`
	AuthoredDate   time.Time `json:"authored_date"`
	CommitterName  string    `json:"committer_name"`
	CommitterEmail string    `json:"committer_email"`
	CommittedDate  time.Time `json:"committed_date"`
	Trailers       struct {
	} `json:"trailers"`
	ExtendedTrailers struct {
	} `json:"extended_trailers"`
	WebURL string `json:"web_url"`
}

type Tasks []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Path string `json:"path"`
	Mode string `json:"mode"`
}

type Branch []struct {
	Name   string `json:"name"`
	Commit struct {
		ID             string    `json:"id"`
		ShortID        string    `json:"short_id"`
		CreatedAt      time.Time `json:"created_at"`
		ParentIds      []string  `json:"parent_ids"`
		Title          string    `json:"title"`
		Message        string    `json:"message"`
		AuthorName     string    `json:"author_name"`
		AuthorEmail    string    `json:"author_email"`
		AuthoredDate   time.Time `json:"authored_date"`
		CommitterName  string    `json:"committer_name"`
		CommitterEmail string    `json:"committer_email"`
		CommittedDate  time.Time `json:"committed_date"`
		Trailers       struct {
		} `json:"trailers"`
		ExtendedTrailers struct {
		} `json:"extended_trailers"`
		WebURL string `json:"web_url"`
	} `json:"commit"`
	Merged             bool   `json:"merged"`
	Protected          bool   `json:"protected"`
	DevelopersCanPush  bool   `json:"developers_can_push"`
	DevelopersCanMerge bool   `json:"developers_can_merge"`
	CanPush            bool   `json:"can_push"`
	Default            bool   `json:"default"`
	WebURL             string `json:"web_url"`
}

// Получение коммитов из репозитория
func GetCommit(projectID string, branch string, path string) Commit {
	// Запрос к GitLab API для получения коммитов
	resp := SendRequest("api/v4/projects/" + projectID + "/repository/commits?ref_name=" + branch + "&path=" + path)
	defer resp.Body.Close()

	// Декодирование ответа
	log.Println("Декодирование коммитов")
	var commit Commit
	if err := json.NewDecoder(resp.Body).Decode(&commit); err != nil {
		log.Printf("\033[31mОшибка декодирования ответа: %v\033[0m", err)
	}
	log.Printf("Получено %d", len(commit))

	return commit
}

// Получение merge requests из репозитория
func GetMergeRequests(projectID string, state string) MergeRequest {
	// Запрос к GitLab API для получения merge requests
	resp := SendRequest("api/v4/projects/" + projectID + "/merge_requests?state=" + state)
	defer resp.Body.Close()

	// Декодирование ответа в срез структур MergeRequest
	log.Println("Декодирование merge requests")
	var mergeRequests MergeRequest
	if err := json.NewDecoder(resp.Body).Decode(&mergeRequests); err != nil {
		log.Printf("\033[31mОшибка декодирования ответа: %v\033[0m", err)
	}
	log.Printf("Получено %d", len(mergeRequests))

	return mergeRequests
}

// Получение YAML-файла
func GetYAML(projectID string, openMR MergeRequest, task string, category string) []byte {
	sourceBranch := openMR[0].SourceBranch
	filePath := fmt.Sprintf("tasks/%s/%s/challenge.yml", category, task)
	escapedPath := url.PathEscape(filePath)
	escapedBranch := url.QueryEscape(sourceBranch)

	// Собираем путь к API
	apiPath := fmt.Sprintf("api/v4/projects/%s/repository/files/%s/raw?ref=%s", projectID, escapedPath, escapedBranch)
	log.Printf("Получение YAML из %s", apiPath)

	// Отправляем запрос к GitLab API
	resp := SendRequest(apiPath)
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("\033[31mОшибка при чтении YAML: %v\033[0m", err)
	}

	log.Printf("YAML для %s получен успешно", task)
	log.Printf("Длина YAML-данных для %s/%s: %d байт", category, task, len(data))

	return data
}

// Получение тасков из репозитория
func GetTasks(projectID string, category string) Tasks {
	// Запрос к GitLab API для получения тасков
	resp := SendRequest("api/v4/projects/" + projectID + "/repository/tree?path=tasks/" + category)
	defer resp.Body.Close()

	// Проверка статуса
	if resp.StatusCode == http.StatusNotFound {
		log.Printf("Категория %s не найдена (404), пропускаем", category)
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("\033[31mНеуспешный ответ от GitLab: %d — %s\033[0m", resp.StatusCode, string(body))
		return nil
	}

	// Парсинг JSON
	var tasks Tasks
	if err := json.NewDecoder(resp.Body).Decode(&tasks); err != nil {
		log.Printf("\033[31mОшибка декодирования тасков для категории %s: %v\033[0m", category, err)
		return nil
	}
	log.Printf("Категория %s: получено %d тасков", category, len(tasks))
	return tasks
}

// Получение веток из репозитория
func GetBranches(projectID string) Branch {
	// Запрос к GitLab API для получения тасков
	resp := SendRequest("api/v4//projects/" + projectID + "/repository/branches")
	defer resp.Body.Close()

	// Декодирование ответа в срез структур Tasks
	log.Println("Декодирование веток")
	var branch Branch
	if err := json.NewDecoder(resp.Body).Decode(&branch); err != nil {
		log.Printf("\033[31mОшибка декодирования ответа\033[0m: %v", err)
	}
	log.Printf("Получено %d", len(branch))

	return branch
}
