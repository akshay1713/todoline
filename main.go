package main

import (
	"bufio"
	"fmt"
	"github.com/akshay1713/todoline/climanager"
	"github.com/urfave/cli"
	"os"
	"strconv"
)

type todolineConfig struct {
	AuthToken string
}

func main() {
	app := cli.NewApp()
	app.Name = "Todoline"
	app.Usage = "Todoist for the command line"
	app.Commands = []cli.Command{
		{
			Name:  "projects",
			Usage: "Handle Todoist Projects",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "Add new projects",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "file, f",
							Usage: "File path for adding multiple projects",
						},
					},
					Action: func(c *cli.Context) error {
						auth_token := getAuthToken()
						cm := climanager.InitCliManager(auth_token)
						project_names := c.Args()
						file := c.String("file")
						if file != "" {
							names_from_file := readLine(file)
							for _, name := range names_from_file {
								project_names = append(project_names, name)
							}
						}
						if len(project_names) == 0 {
							fmt.Println("Please enter at least one project name")
							return nil
						}
						fmt.Println("Adding Projects")
						cm.AddProjects(project_names)
						return nil
					},
				},
				{
					Name:  "list",
					Usage: "List All Projects",
					Action: func(c *cli.Context) error {
						auth_token := getAuthToken()
						cm := climanager.InitCliManager(auth_token)
						cm.ListProjects()
						return nil
					},
				},
				{
					Name:  "expand",
					Usage: "Expand the given project",
					Action: func(c *cli.Context) error {
						auth_token := getAuthToken()
						cm := climanager.InitCliManager(auth_token)
						project_id, parse_err := strconv.ParseInt(c.Args().Get(0), 10, 64)
						if parse_err != nil {
							fmt.Println("Please enter a valid project id")
							return nil
						}
						cm.ListItemsForProject(project_id)
						return nil
					},
				},
				{
					Name:  "delete",
					Usage: "Delete the given project",
					Action: func(c *cli.Context) error {
						auth_token := getAuthToken()
						cm := climanager.InitCliManager(auth_token)
						args := c.Args()
						var project_ids []int64
						var id int64
						var parse_err error
						for _, id_string := range args {
							id, parse_err = strconv.ParseInt(id_string, 10, 64)
							if parse_err != nil {
								fmt.Println(id_string + " is not a valid id. Please enter valid project ids")
							}
							project_ids = append(project_ids, id)
						}
						cm.DeleteProjects(project_ids)
						return nil
					},
				},
				{
					Name:  "share",
					Usage: "Share the given project with the given email",
					Action: func(c *cli.Context) error {
						auth_token := getAuthToken()
						cm := climanager.InitCliManager(auth_token)
						args := c.Args()
						email := args.Get(0)
						project_id, parse_err := strconv.ParseInt(args.Get(1), 10, 64)
						if parse_err != nil {
							fmt.Println(args.Get(1) + " is not a valid project id.")
						}
						cm.ShareProject(email, project_id)
						return nil
					},
				},
				{
					Name:  "unshare",
					Usage: "Remove the given collaborator from the given project",
					Action: func(c *cli.Context) error {
						auth_token := getAuthToken()
						cm := climanager.InitCliManager(auth_token)
						args := c.Args()
						email := args.Get(0)
						project_id, parse_err := strconv.ParseInt(args.Get(1), 10, 64)
						if parse_err != nil {
							fmt.Println(args.Get(1) + " is not a valid project id.")
						}
						cm.UnshareProject(email, project_id)
						return nil
					},
				},
				{
					Name:  "export",
					Usage: "Export all projects to a json file",
					Action: func(c *cli.Context) error {
						auth_token := getAuthToken()
						cm := climanager.InitCliManager(auth_token)
						cm.ExportProjectItems()
						return nil
					},
				},
			},
		},
		{
			Name:  "items",
			Usage: "Handle Todoist items",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "Add new items",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "file, f",
							Usage: "File path for adding multiple items.",
						},
					},
					Action: func(c *cli.Context) error {
						auth_token := getAuthToken()
						cm := climanager.InitCliManager(auth_token)
						args := c.Args()
						project_id_string := args[len(args)-1]
						project_id, parse_err := strconv.ParseInt(project_id_string, 10, 64)
						if parse_err != nil {
							fmt.Println("Please enter a valid project id")
							return nil
						}
						item_names := args[:len(args)-1]
						file := c.String("file")
						if file != "" {
							names_from_file := readLine(file)
							for _, name := range names_from_file {
								item_names = append(item_names, name)
							}
						}
						if len(item_names) == 0 {
							fmt.Println("Please enter at least one item name")
							return nil
						}
						fmt.Println("Adding Items")
						cm.AddItems(item_names, project_id)
						return nil
					},
				},
				{
					Name:  "list",
					Usage: "List All Items",
					Action: func(c *cli.Context) error {
						auth_token := getAuthToken()
						cm := climanager.InitCliManager(auth_token)
						args := c.Args()
						if len(args) > 0 {
							project_id, parse_err := strconv.ParseInt(c.Args().Get(0), 10, 64)
							if parse_err != nil {
								fmt.Println("Please enter a valid project id")
								return nil
							}
							cm.ListItemsForProject(project_id)
						} else {
							cm.ListItems()
						}
						return nil
					},
				},
				{
					Name:  "complete",
					Usage: "Complete the given item",
					Action: func(c *cli.Context) error {
						auth_token := getAuthToken()
						cm := climanager.InitCliManager(auth_token)
						args := c.Args()
						var item_ids []int64
						var item_id int64
						var parse_err error
						for _, arg := range args {
							item_id, parse_err = strconv.ParseInt(arg, 10, 64)
							if parse_err != nil {
								fmt.Println(arg + " is not a valid value. Please enter a valid item id")
								return nil
							}
							item_ids = append(item_ids, item_id)
						}
						cm.CompleteItems(item_ids)
						return nil
					},
				},
				{
					Name:  "delete",
					Usage: "Delete the given item",
					Action: func(c *cli.Context) error {
						auth_token := getAuthToken()
						cm := climanager.InitCliManager(auth_token)
						args := c.Args()
						var item_ids []int64
						var id int64
						var parse_err error
						for _, id_string := range args {
							id, parse_err = strconv.ParseInt(id_string, 10, 64)
							if parse_err != nil {
								fmt.Println(id_string + " is not a valid id. Please enter valid project ids")
							}
							item_ids = append(item_ids, id)
						}
						cm.DeleteItems(item_ids)
						return nil
					},
				},
			},
		},
		{
			Name:  "setup",
			Usage: "Setup Todoline",
			Action: func(c *cli.Context) error {
				auth_token := c.Args().Get(0)
				fmt.Println("Creating config file and storing the token in it")
				setupConfig(auth_token)
				//fmt.Println("Getting inbox id and storing it")
				//cm := climanager.InitCliManager(auth_token)
				//inbox_id := cm.GetInboxId()
				//if inbox_id > 0 {
				//saveInboxId(inbox_id)
				//}
				return nil
			},
		},
	}
	app.Run(os.Args)
}

func getAuthToken() string {
	config := getConfig()
	auth_token := config.Get("auth_token")
	if auth_token == nil {
		fmt.Println("Auth token not found in config file. Please run setup and provide your auth token")
		return ""
	}
	return auth_token.(string)
}

func readLine(path string) []string {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	var names []string
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}
	return names
}
