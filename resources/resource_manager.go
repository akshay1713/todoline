package resources

import (
	"github.com/akshaysingh1713/gotodoist"
)

func getTodoistAPI(auth_token string) gotodoist.TodoistAPI {
	todoist_api := gotodoist.InitTodoistAPI(auth_token)
	return todoist_api
}

type Resources struct {
	AuthToken  string
	todoistAPI gotodoist.TodoistAPI
}

func InitResources(auth_token string) Resources {
	todoist_api := gotodoist.InitTodoistAPI(auth_token)
	resources := Resources{todoistAPI: todoist_api}
	return resources
}
