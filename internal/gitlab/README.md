# Пакет gitlab
Предназначен для работы с GitLab REST API: создания сервиса, отправки запросов и парсинга YAML-файлов

## `func InitClient()`
Инициализирует новго клиента для запросов к API. Необходим токен GitLab с разрешениями `read_api`, `read_repository`. 
## `func NewRequest(method, path string) (*http.Request, error)`
Создаёт новый запрос к GitLab API. На вход принимает метод (в нашем случае в основном GET) и [эндпоинт API](https://docs.gitlab.com/api/rest/). Возвращает указатель на объект типа [http.Request](https://pkg.go.dev/net/http#Request) - HTTP-запрос, и объект типа error - ошибку.
## `func SendRequest(path string) (*http.Response, error)`
На основе инициализированного клиента остправляет созданный запрос. На вход принимает [эндпоинт API](https://docs.gitlab.com/api/rest/). Возвращает указатель на объект типа [http.Response](https://pkg.go.dev/net/http#Response) - HTTP-ответ и объект типа error - ошибку.

## `type MergeRequest []struct`
Срез структуры необходимый для парсинга JSON, возвращаемого эндпоинтом [api/v4/projects/:projectID/merge_requests?state=:state](https://docs.gitlab.com/api/merge_requests/).
## `type Commit []struct`
Срез структуры необходимый для парсинга JSON, возвращаемая эндпоинтом [api/v4/projects/:projectID/repository/commits?ref_name=:branch&path=:path](https://docs.gitlab.com/api/commits/).
## `type Tasks []struct`
Срез структуры необходимый для парсинга JSON, возвращаемая эндпоинтом [api/v4/projects/2/repository/tree?path=:path](https://docs.gitlab.com/api/repositories/).

## `func GetMergeRequests(projectID string, state string) (MergeRequest, error)`
Предназначена для отправки запроса на эндпоинт [api/v4/projects/:projectID/merge_requests?state=:state](https://docs.gitlab.com/api/merge_requests/) и декодирования ответа в экземпляр структуры MergeRequest. Принимает на вход ID проекта, и [статус Merge Request](https://docs.gitlab.com/api/merge_requests/). Возвращает экземпляр структуры MergeRequest и объект типа error - ошибку.
### Возможные статусы Merge Request
- `opened` — только открытые
- `merged` — только смерженные
- `closed` — закрытые, но не смерженные
- `all` — все (по умолчанию)
## `func GetCommit(projectID string, branch string, path string) (Commit, error)`
Предназначена для отправки запроса на эндпоинт [api/v4/projects/:projectID/repository/commits?ref_name=:branch&path=:path](https://docs.gitlab.com/api/commits/) и декодирования ответа в экземпляр структуры Commit. Принимает на вход ID проекта, ветку и путь проекта. Возвращает экземпляр структуры Commit и объект типа error - ошибку.
## `func GetYAML(projectID string, openMR MergeRequest, task string, category string) ([]byte, error)`
Предназначена для получения файла challenge.yml из Merge Request и декодирования ответа в массив байт. Принимает на вход ID проекта, экземпляр структуры MergeRequest, хранящий открытые Merge Request, название таска и категорию. Возвращает массив байт и объект типа error - ошибку.
## `func GetTasks(projectID string, category string) (Tasks, error)`
Предназначена для отправки запроса на эндпоинт [api/v4/projects/2/repository/tree?path=:path](https://docs.gitlab.com/api/repositories/) и декодирования ответа в экземпляр структуры Tasks. Принимает на вход ID проекта и категорию таска. Возвращает экземпляр структуры Tasks и объект типа error - ошибку.

## `type Task struct`
Структура необходимая для парсинга YAML-файла, получаемого из Merge Request с эндпоинта `api/v4/projects/:projectID/repository/files/:path/raw?ref=:branch`.
## `func ParseTask(data []byte) (Task, error)`
Пердназначена для парсинга YAML-файла и декодирования его в экземпляр структуры Task. Принимает на вход массив байт. Возвращает экземпляр структуры Task и объект типа error - ошибку.