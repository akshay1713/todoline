package main

import (
	"fmt"
	"github.com/akshaysingh1713/todoline/resources"
)

func listProjects() {
	config := getConfig()
	auth_token := config.Get("auth_token").(string)
	resources := resources.InitResources(auth_token)
	projects, err := resources.GetAllProjects()
	for _, project := range projects {
		fmt.Printf("%s -  %v\n", project["name"], int(project["id"].(float64)))
	}
	if err != nil {
		fmt.Println(err)
	}
}

func addProjects(project_names []string) {
	config := getConfig()
	auth_token := config.Get("auth_token").(string)
	resources := resources.InitResources(auth_token)
	if len(project_names) == 0 {
		fmt.Println("at least one project name is required")
	}
	response, err := resources.AddProject(project_names)
	if err != nil {
		fmt.Println(err)
	} else if response["status"] == "200 OK" {
		fmt.Println("Added projects successfully")
	} else {
		fmt.Println("Unexpected response " + response["status"].(string))
	}
}
