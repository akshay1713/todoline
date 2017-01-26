package main

import (
	"fmt"
	"github.com/akshaysingh1713/todoline/climanager"
	"github.com/urfave/cli"
	"os"
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
					Action: func(c *cli.Context) error {
						auth_token := getAuthToken()
						cm := climanager.InitCliManager(auth_token)
						project_names := c.Args()
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
			},
		},
		{
			Name:  "setup",
			Usage: "Setup Todoline",
			Action: func(c *cli.Context) error {
				auth_token := c.Args().Get(0)
				setupConfig(auth_token)
				return nil
			},
		},
	}
	app.Run(os.Args)
}

func getAuthToken() string {
	config := getConfig()
	auth_token := config.Get("auth_token").(string)
	return auth_token
}
