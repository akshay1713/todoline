package climanager

import (
	"encoding/json"
	"fmt"
	"strconv"
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
	name_ids := response["name_ids"].(map[string]int64)
	for k, v := range name_ids {
		fmt.Println(k + " " + strconv.FormatInt(v, 10))
	}
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

func (cm CliManager) ShareProject(email string, project_id int64) {
	response, err := cm.resources.ShareProject(email, project_id)
	if err != nil {
		fmt.Println("Error while sharing project:\n")
		fmt.Println(err)
	} else if response["status"] == "200 OK" {
		fmt.Println("Shared project successfully")
	}
}

func (cm CliManager) UnshareProject(email string, project_id int64) {
	response, err := cm.resources.UnshareProject(email, project_id)
	if err != nil {
		fmt.Println("Error while unsharing project:\n")
		fmt.Println(err)
	} else if response["status"] == "200 OK" {
		fmt.Println("Removed project collaborator successfully")
	}
}

func (cm CliManager) ExportProjectItems() {
	projects, err := cm.resources.GetAllProjects()
	if err != nil {
		fmt.Println("Error while getting projects")
		fmt.Println(err)
	}
	project_items := map[string][]string{}
	project_id_names := map[int64]string{}
	project_id_names[90] = "rest"
	var project_id int64
	for _, project := range projects {
		project_id = int64(project["id"].(float64))
		project_id_names[project_id] = project["name"].(string)
		project_items[project["name"].(string)] = []string{}
	}
	items, err := cm.resources.GetAllItems()
	if err != nil {
		fmt.Println("Error while getting items")
		fmt.Println(err)
	}

	for _, item := range items {
		project_id = int64(item["project_id"].(float64))
		project_name := project_id_names[project_id]
		project_items[project_name] = append(project_items[project_name], item["content"].(string))
	}
	project_items_json, marshal_err := json.MarshalIndent(project_items, "", " ")
	if marshal_err != nil {
		fmt.Println(marshal_err)
	}
	fmt.Println(string(project_items_json))
}
