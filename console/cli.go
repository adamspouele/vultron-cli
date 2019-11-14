package console

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

const (
	do_token_env_var = "VULTRON_DO_TOKEN"
)

func Handle() {

	fmt.Printf(os.Getenv(do_token_env_var))

	app := cli.NewApp()
	app.Name = "Vultron Cli"
	app.Usage = "Vultron is a Paas based on nomand."

	app.Commands = []*cli.Command{
		{
			Name:     "add",
			Category: "Basic Commands",
			Aliases:  []string{"a"},
			Usage:    "add a task to the list",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:     "complete",
			Category: "Deploy Commands",
			Aliases:  []string{"c"},
			Usage:    "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:     "complete2",
			Category: "Cluster Management Commands",
			Aliases:  []string{"c"},
			Usage:    "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:     "complete3",
			Category: "Troubleshooting and Debugging Commands",
			Aliases:  []string{"c"},
			Usage:    "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:     "complete4",
			Category: "Advanced Commands",
			Aliases:  []string{"c"},
			Usage:    "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:     "complete5",
			Category: "Settings Commands",
			Aliases:  []string{"c"},
			Usage:    "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "options for task templates",
			Subcommands: []*cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
