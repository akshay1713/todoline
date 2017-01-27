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
		return
	}
}

func (cm CliManager) AddProjects(project_names []string) {
	response, err := cm.resources.AddProject(project_names)
	if err != nil {
		fmt.Println(err)
	} else if response["status"] == "200 OK" {
		fmt.Println("Added projects successfully")
	} else {
		fmt.Println("Unexpected response " + response["status"].(string))
	}
}

func (cm CliManager) ExpandProject(project_id int64) {
	projects, err := cm.resources.GetAllProjects()
	for _, project := range projects {
		if int64(project["id"].(float64)) == project_id {
			fmt.Printf("%s -  %v\n", project["name"], int(project["id"].(float64)))
			fmt.Println(project)
			return
		}
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("No project found with the given id")
}

func (cm CliManager) GetInboxId() int64 {
	projects, err := cm.resources.GetAllProjects()
	if err != nil {
		fmt.Println("Error while getting projects:\n")
		fmt.Println(err)
	}
	for _, project := range projects {
		if project["name"] == "Inbox" {
			return int64(project["id"].(float64))
		}
	}
	return -1
}

func (cm CliManager) DeleteProjects(project_ids []int64) {
	response, err := cm.resources.DeleteProjects(project_ids)
	if err != nil {
		fmt.Println("Error while deleting projects:\n")
		fmt.Println(err)
	} else if response["status"] == "200 OK" {
		fmt.Println("Deleted project(s) successfully")
	}
}
