package cloud

import (
	"fmt"
	"log"
	"strconv"

	consul "github.com/adamspouele/vultron-cli/consul/script"
	"github.com/adamspouele/vultron-cli/naming/label"
	"github.com/adamspouele/vultron-cli/naming/tag"
	"github.com/digitalocean/godo"
)

// CreateCluster create a cluster in the cloud
func CreateCluster(name string, region string, serverSize int, clientSize int, sshKey godo.DropletCreateSSHKey) {

	fmt.Printf("> Plan deployment of %v server nodes\n ", serverSize)

	fmt.Printf("> Start creating %v server nodes in %v cluster\n ", clientSize, name)

	for i := 0; i < serverSize; i++ {
		newDroplet, err := nodeCreationProcess(name, NodeKindCluster, NodeResServer, i, region, sshKey, consul.GetConsulServerInstallScript())

		if err != nil {
			log.Fatalln("! Error creating server node %v", i)
		} else {
			fmt.Printf("	+ Node %v successfully created\n", newDroplet.Name)
		}
	}

	fmt.Printf("> Start creating %v client nodes in cluster %v\n ", clientSize, name)

	for k := 0; k < clientSize; k++ {
		newDroplet, err := nodeCreationProcess(name, NodeKindCluster, NodeResClient, k, region, sshKey, consul.GetConsulClientInstallScript())

		if err != nil {
			log.Fatalln("! Error creating client node %v", k)
		} else {
			fmt.Printf("	+ Node %v successfully created.\n ", newDroplet.Name)
		}
	}

	fmt.Println("> Cluster Created.")
}

/*
nodeCreationProcess create a node in a cluster
NodeKind can be of 2 value : [standalone, cluster]
NodeRes can be of 2 value : [server, client]
*/
func nodeCreationProcess(clusterName string, nodeKind NodeKind, nodeRes NodeRes, iteration int, region string, sshKey godo.DropletCreateSSHKey, userData string) (*godo.Droplet, error) {
	nodeName := label.GenerateClientLabel(label.Label(clusterName), iteration)

	nodeTags := getNodeTags(clusterName, getNodeKind(nodeKind), getNodeRes(nodeRes))

	// "#!/bin/bash \n  mkdir -p /etc/vultron;\n  cat << 'Vultron ecosystem' > /etc/vultron/README.txt"

	return CreateNode(nodeName+"-"+string(nodeRes)+strconv.Itoa(iteration), region, sshKey, userData, nodeTags)
}

func getNodeKind(nodeKind NodeKind) NodeKind {
	if nodeKind == "cluster" {
		return NodeKind(tag.GetClusterKindTag())
	}

	return NodeKind(tag.GetStandaloneKindTag())
}

func getNodeRes(nodeRes NodeRes) NodeRes {
	if nodeRes == "server" {
		return NodeRes(tag.GetServerResourceTag())
	}

	return NodeRes(tag.GetClientResourceTag())
}

func getNodeTags(clusterName string, nodeKind NodeKind, nodeRes NodeRes) []string {
	return []string{
		string(nodeKind),
		string(nodeRes),
		tag.GetPropTag("cluster-name", tag.TagProp(clusterName)),
	}
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

	fmt.Printf("> Deleting cluster %v...\n", clusterTag)

	client, ctx, _ := GetDoClient()

	nodesTag := tag.GetPropTag(tag.TagProp("cluster-name"), tag.TagProp(clusterTag))

	response, err := client.Droplets.DeleteByTag(ctx, nodesTag)

	if err != nil {
		log.Fatalln("! Error deleting cluster [%v] : %v\n ", clusterTag, err)
	}

	fmt.Println("\n Done.")

	return response, err
}
