package climanager

import (
	"fmt"
	"strconv"
)

func (cm CliManager) ListItems() {
	items, err := cm.resources.GetAllItems()
	if err != nil {
		fmt.Println(err)
		return
	}
	printItems(items)
}

func (cm CliManager) ListItemsForProject(project_id int64) {
	items, err := cm.resources.GetAllItems()
	if err != nil {
		fmt.Println(err)
		return
	}
	var project_items []map[string]interface{}
	for _, item := range items {
		if int64(item["project_id"].(float64)) == project_id {
			project_items = append(project_items, item)
		}
	}
	printItems(project_items)

}

func (cm CliManager) AddItems(item_names []string, project_id int64) {
	response, err := cm.resources.AddItem(item_names, project_id)
	name_ids := response["name_ids"].(map[string]int64)
	for k, v := range name_ids {
		fmt.Println(k + " " + strconv.FormatInt(v, 10))
	}
	if err != nil {
		fmt.Println(err)
	} else if response["status"] == "200 OK" {
		fmt.Println("Added items successfully")
	} else {
		fmt.Println("Unexpected response " + response["status"].(string))
	}
}

func (cm CliManager) CompleteItems(item_ids []int64) {
	response, err := cm.resources.CompleteItems(item_ids)
	if err != nil {
		fmt.Println("Error occured:\n%v", response)
	}
}

func printItems(items []map[string]interface{}) {
	var indent, checked int
	var content string
	closed := "\u2713"
	open := "\u2717"
	for _, item := range items {
		content = ""
		indent = int(item["indent"].(float64))
		for i := 1; i < indent; i++ {
			content += " "
		}
		checked = int(item["checked"].(float64))
		if checked == 1 {
			content += closed
		} else {
			content += open
		}
		content += item["content"].(string)
		fmt.Printf("%s -  %v\n", content, int(item["id"].(float64)))
	}
}

func (cm CliManager) DeleteItems(item_ids []int64) {
	response, err := cm.resources.DeleteItems(item_ids)
	if err != nil {
		fmt.Println("Error while deleting items:\n")
		fmt.Println(err)
	} else if response["status"] == "200 OK" {
		fmt.Println("Deleted item(s) successfully")
	}
}
