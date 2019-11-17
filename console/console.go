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

	fmt.Println(os.Getenv(do_token_env_var))

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
					Name:    "cluster",
					Usage:   "create a cluster",
					Aliases: []string{"cl"},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "name",
							Value:   "",
							Aliases: []string{"n"},
							Usage:   "The cluster name",
						},
						&cli.StringFlag{
							Name:    "region",
							Value:   "ams3",
							Aliases: []string{"r"},
							Usage:   "The cluster's region",
						},
						&cli.IntFlag{
							Name:    "consul-server-size",
							Value:   1,
							Aliases: []string{"css"},
							Usage:   "The consul server nodes size",
						},
						&cli.IntFlag{
							Name:    "nomad-server-size",
							Value:   1,
							Aliases: []string{"nss"},
							Usage:   "The nomad server nodes size",
						},
						&cli.IntFlag{
							Name:    "client-size",
							Value:   1,
							Aliases: []string{"cs"},
							Usage:   "The client nodes size",
						},
						&cli.StringFlag{
							Name:    "sshKey",
							Value:   "",
							Aliases: []string{"sh"},
							Usage:   "The ssh Fingerprint",
						},
					},
					Action: func(c *cli.Context) error {

						if len(c.String("name")) > 4 && len(c.String("sshKey")) > 10 {
							sshKey := godo.DropletCreateSSHKey{
								Fingerprint: c.String("sshKey"),
							}

							cloud.CreateCluster(c.String("name"),
								c.String("region"),
								c.Int("consul-server-size"),
								c.Int("nomad-server-size"),
								c.Int("client-size"),
								sshKey,
							)
						} else {
							log.Fatalln("You must provide at least the name of the cluster which have a length superior to 4 and an sshKey.")
						}

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
					Name:    "cluster",
					Usage:   "delete a cluster",
					Aliases: []string{"cl"},
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:    "name",
							Value:   "",
							Aliases: []string{"n"},
							Usage:   "The cluster name",
						},
					},
					Action: func(c *cli.Context) error {
						name := c.String("name")

						if len(name) > 0 {
							cloud.DeleteCluster(name)
						} else {
							log.Fatalln("! You must provide a cluster name with option 'name'.")
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
