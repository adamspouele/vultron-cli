package cloud

import (
	"fmt"

	"github.com/digitalocean/godo"
)

// reference : https://developers.digitalocean.com/documentation/changelog/api-v2/new-size-slugs-for-droplet-plan-changes/
// NodeSize is the different node sizes available
var NodeSize = map[string]string{
	"small-xl":  "s-1vcpu-1gb",  // s-1vcpu-1gb
	"small-x":   "s-1vcpu-2gb",  // s-1vcpu-2gb
	"small":     "s-2vcpu-2gb",  // s-2vcpu-2gb
	"medium":    "s-2vcpu-4gb",  // s-2vcpu-4gb
	"medium-x":  "s-4vcpu-8gb",  // s-4vcpu-8gb
	"medium-xl": "s-6vcpu-16gb", // s-6vcpu-16gb
	"large":     "s-8vcpu-32gb", // s-8vcpu-32gb
	"large-x":   "s-8vcpu-32gb", // s-12vcpu-48gb
	"large-xl":  "s-8vcpu-32gb", // s-16vcpu-64gb
}

// NodeKind define the kind of node
type NodeKind string

const (
	NodeKindStandalone NodeKind = "standalone"
	NodeKindCluster    NodeKind = "cluster"
)

// NodeRes define the type of resource of a node
type NodeRes string

const (
	NodeResServer NodeRes = "server"
	NodeResClient NodeRes = "client"
)

// CreateNode create a new node with specific configuration
func CreateNode(name string, region string, sshKey godo.DropletCreateSSHKey, userData string, tags []string) (*godo.Droplet, error) {
	client, doContext, _ := GetDoClient()

	createRequest := &godo.DropletCreateRequest{
		Name:   name,
		Region: region,
		Size:   "s-1vcpu-1gb",
		Image: godo.DropletCreateImage{
			Slug: "ubuntu-16-04-x64",
		},
		SSHKeys: []godo.DropletCreateSSHKey{
			sshKey,
		},
		IPv6:              true,
		PrivateNetworking: true,
		UserData:          userData,
		Monitoring:        true,
		Tags:              tags,
	}

	newDroplet, _, err := client.Droplets.Create(doContext, createRequest)

	if err != nil {
		fmt.Println("! Something bad happened: %s\n\n", err)
	}

	return newDroplet, err
}
