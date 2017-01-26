package climanager

import (
	"fmt"
)

func (cm CliManager) ListItems() {
	items, err := cm.resources.GetAllItems()
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
	if err != nil {
		fmt.Println(err)
	}
}

func (cm CliManager) AddItems(item_names []string, project_id int64) {
	if len(item_names) == 0 {
		fmt.Println("at least one item name is required")
	}
	response, err := cm.resources.AddItem(item_names, project_id)
	fmt.Println(response)
	if err != nil {
		fmt.Println(err)
	} else if response["status"] == "200 OK" {
		fmt.Println("Added items successfully")
	} else {
		fmt.Println("Unexpected response " + response["status"].(string))
	}
}
