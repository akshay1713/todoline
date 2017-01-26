package resources

func (resources Resources) GetAllItems() ([]map[string]interface{}, error) {
	response, err := resources.todoistAPI.Items.GetAll()
	resp_body := response["body"].(map[string]interface{})
	projects_interface := resp_body["items"].([]interface{})
	projects := make([]map[string]interface{}, len(projects_interface))
	for i, project := range projects_interface {
		projects[i] = project.(map[string]interface{})
	}
	return projects, err
}

func (resources Resources) AddItem(item_names []string, project_id int64) (map[string]interface{}, error) {
	response, err := resources.todoistAPI.Items.Add(item_names, project_id)
	return response, err
}
