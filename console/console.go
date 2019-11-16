package console

import (
	"fmt"
	"log"
	"os"

	"github.com/adamspouele/vultron-cli/cloud"
	"github.com/digitalocean/godo"
	"github.com/urfave/cli"
)

const (
	do_token_env_var = "VULTRON_DO_TOKEN"
)

// Handle execute the CLI
func Handle() {

	fmt.Printf(os.Getenv(do_token_env_var))

	app := cli.NewApp()
	app.Name = "Vultron Cli"
	app.Usage = "Vultron is a Paas based on nomand."

	app.Commands = []*cli.Command{
		{
			Name:     "explain",
			Category: "Basic Commands",
			Aliases:  []string{"ex"},
			Usage:    "Documentation of resources",
			Subcommands: []*cli.Command{
				{
					Name:  "cluster",
					Usage: "explain a cluster",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Value: "",
							Usage: "The cluster name",
						},
					},
					Action: func(c *cli.Context) error {
						name := c.String("name")

						if len(name) > 0 {
							cloud.ExplainCluster(name)
						} else {
							log.Fatalln("! You must provide a cluster name with option 'name' to get a explaination.")
						}

						return nil
					},
				},
			},
		},
		{
			Name:     "create",
			Category: "Basic Commands",
			Aliases:  []string{"ct"},
			Usage: "Create	a resource from a file or interactively",
			Subcommands: []*cli.Command{
				{
					Name:  "cluster",
					Usage: "create a cluster",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Value: "",
							Usage: "The cluster name",
						},
						&cli.IntFlag{
							Name:  "serverSize",
							Value: 0,
							Usage: "The server node size",
						},
						&cli.IntFlag{
							Name:  "clientSize",
							Value: 0,
							Usage: "The client node size",
						},
						&cli.StringFlag{
							Name:  "sshKey",
							Value: "",
							Usage: "The ssh Fingerprint",
						},
					},
					Action: func(c *cli.Context) error {

						sshKey := godo.DropletCreateSSHKey{
							Fingerprint: c.String("sshKey"),
						}

						cloud.CreateCluster(c.String("name"),
							c.String("region"),
							c.Int("serverSize"),
							c.Int("clientSize"),
							sshKey,
						)

						return nil
					},
				},
			},
		},
		{
			Name:     "delete",
			Category: "Basic Commands",
			Aliases:  []string{"del"},
			Usage:    "Delete resources from a file or interactively",
			Subcommands: []*cli.Command{
				{
					Name:  "cluster",
					Usage: "delete a cluster",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Value: "",
							Usage: "The cluster name",
						},
					},
					Action: func(c *cli.Context) error {
						name := c.String("name")

						if len(name) > 0 {
							cloud.ExplainCluster(name)
						} else {
							log.Fatalln("! You must provide a cluster name with option 'name' to get a explaination.")
						}

						return nil
					},
				},
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
