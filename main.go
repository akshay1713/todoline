package main

import (
	"fmt"
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
						project_names := c.Args()
						fmt.Println("Adding Projects")
						addProjects(project_names)
						return nil
					},
				},
				{
					Name:  "list",
					Usage: "List All Projects",
					Action: func(c *cli.Context) error {
						listProjects()
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
