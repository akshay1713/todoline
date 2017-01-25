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

func (resources Resources) GetAllProjects() ([]map[string]interface{}, error) {
	response, err := resources.todoistAPI.Projects.GetAll()
	resp_body := response["body"].(map[string]interface{})
	projects_interface := resp_body["projects"].([]interface{})
	projects := make([]map[string]interface{}, len(projects_interface))
	for i, project := range projects_interface {
		projects[i] = project.(map[string]interface{})
	}
	return projects, err
}

func (resources Resources) AddProject(project_names []string) (map[string]interface{}, error) {
	response, err := resources.todoistAPI.Projects.Add(project_names)
	return response, err
}
