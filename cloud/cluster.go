package cloud

import (
	"fmt"
	"log"

	"github.com/adamspouele/vultron-cli/cloud"

	"github.com/adamspouele/vultron-cli/naming/label"
	"github.com/adamspouele/vultron-cli/naming/tag"
	"github.com/digitalocean/godo"
)

// CreateCluster create a cluster in the cloud
func CreateCluster(name string, region string, serverSize int, clientSize int, sshKey godo.DropletCreateSSHKey) {

	fmt.Println("> Plan deployment of %v server nodes", serverSize)

	fmt.Println("> Start creating %v server nodes in %v cluster", clientSize, name)

	for i := 0; i < serverSize; i++ {
		newDroplet, err := nodeCreationProcess(name, , cloud.NodeKindCluster, cloud.NodeResServer, i, region, sshKey)

		if err != nil {
			log.Panicln("! Error creating server node %v", i)
		} else {
			fmt.Println("	> Node %v successfully created.", newDroplet.Name)
		}
	}

	fmt.Println("> Start creating %v client nodes in cluster %v", clientSize, name)

	for k := 0; k < clientSize; k++ {
		newDroplet, err := nodeCreationProcess(name, cloud.NodeKindCluster, cloud.NodeResClient, k, region, sshKey)

		if err != nil {
			log.Panicln("! Error creating client node %v", k)
		} else {
			fmt.Println("	> Node %v successfully created.", newDroplet.Name)
		}
	}

	fmt.Println("> Cluster Created.")
}

/*
nodeCreationProcess create a node in a cluster
NodeKind can be of 2 value : [standalone, cluster]
NodeRes can be of 2 value : [server, client]
*/
func nodeCreationProcess(clusterName string, nodeKind NodeKind, nodeRes NodeRes, iteration int, region string, sshKey godo.DropletCreateSSHKey) (*godo.Droplet, error) {
	nodeName := label.GenerateClientLabel(label.Label(clusterName), iteration)

	var currentNodeKind NodeKind
	if nodeKind == "cluster" {
		currentNodeKind := tag.GetClusterKindTag()
	} else {
		currentNodeKind := tag.GetStandaloneKindTag()
	}

	var currentNodeRes NodeRes
	if nodeKind == "server" {
		currentNodeRes := tag.GetServerResourceTag()
	} else {
		currentNodeRes := tag.GetClientResourceTag()
	}

	nodeTags := []string{
		string(currentNodeKind),
		string(currentNodeRes),
		tag.GetPropTag("cluster-name", "clusterName"),
	}

	return CreateNode(nodeName, region, sshKey, "#!/bin/bash cat << 'Vultron ecosystem' > /etc/vultron/README.txt", nodeTags)
}

// ListClusterNodes list cluster nodes
func ListClusterNodes(clusterTag string) ([]godo.Droplet, error) {
	client, ctx, err := GetDoClient()
	opt := &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	}

	if err != nil {
		log.Fatalln("! Error authenticating")
	}

	droplets, _, err := client.Droplets.ListByTag(ctx, clusterTag, opt)

	return droplets, err
}

// ExplainCluster explain a cluster
func ExplainCluster(clusterTag string) {

	droplets, err := ListClusterNodes(clusterTag)

	if err != nil {
		log.Fatalln("! Error while explaining cluster %v", clusterTag)
	}

	for i := 0; i < len(droplets); i++ {
		fmt.Println("# %v : ID='%v' status='%v' name='%v' region='%v' memory='%v' vcpus='%v' image='%v' disk='%v' size='%v' created='%v'",
			i, droplets[i].ID, droplets[i].Status, droplets[i].Name, droplets[i].Region.Slug, droplets[i].Memory,
			droplets[i].Vcpus, droplets[i].Image.Slug, droplets[i].Disk, droplets[i].Size.Slug, droplets[i].Created)
	}
}

// DeleteCluster delete a cluster by removing all cluster's nodes
func DeleteCluster(clusterTag string) (*godo.Response, error) {

	fmt.Println("> Delete cluster %v", clusterTag)

	client, ctx, _ := GetDoClient()

	response, err := client.Droplets.DeleteByTag(ctx, clusterTag)

	if err != nil {
		log.Fatalln("! Error deleting cluster [%v] : %v", clusterTag, err)
	}

	fmt.Println("Done.")

	return response, err
}
