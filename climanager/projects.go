package climanager

import (
	"fmt"
)

func (cm CliManager) ListProjects() {
	projects, err := cm.resources.GetAllProjects()
	for _, project := range projects {
		fmt.Printf("%s -  %v\n", project["name"], int(project["id"].(float64)))
	}
	if err != nil {
		fmt.Println(err)
	}
}

func (cm CliManager) AddProjects(project_names []string) {
	if len(project_names) == 0 {
		fmt.Println("at least one project name is required")
	}
	response, err := cm.resources.AddProject(project_names)
	if err != nil {
		fmt.Println(err)
	} else if response["status"] == "200 OK" {
		fmt.Println("Added projects successfully")
	} else {
		fmt.Println("Unexpected response " + response["status"].(string))
	}
}
