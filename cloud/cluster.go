package cloud

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/adamspouele/vultron-cli/naming/label"
	"github.com/adamspouele/vultron-cli/naming/tag"
	consul "github.com/adamspouele/vultron-cli/package/consul/script"
	nomad "github.com/adamspouele/vultron-cli/package/nomad/script"
	"github.com/digitalocean/godo"
	"github.com/google/uuid"
)

// CreateCluster create a cluster in the cloud
func CreateCluster(name string, region string, consulServerSize int, nomadServerSize int, clientSize int, sshKey godo.DropletCreateSSHKey) {

	nodesCount := consulServerSize + nomadServerSize + clientSize
	fmt.Printf("> Plan deployment of %v nodes\n ", nodesCount)

	// the cluster ID
	clusterID := generateUniqueClusterID()

	fmt.Printf("> Start creating %v consul server nodes in %v cluster\n ", clientSize, name)

	// generate 16 bytes key
	encryptKey := generateClusterEncryptKey()

	for i := 0; i < consulServerSize; i++ {
		newDroplet, err := nodeCreationProcess(name, clusterID, encryptKey, NodeKindConsul, NodeResServer, consulServerSize, i, region, sshKey)

		if err != nil {
			log.Fatalln("! Error creating consul server node %v", i)
		} else {
			fmt.Printf("	+ consul node %v successfully created\n", newDroplet.Name)
		}
	}

	fmt.Printf("> Start creating %v nomad server nodes in %v cluster\n ", clientSize, name)

	for i := 0; i < nomadServerSize; i++ {
		newDroplet, err := nodeCreationProcess(name, clusterID, encryptKey, NodeKindNomad, NodeResServer, nomadServerSize, i, region, sshKey)

		if err != nil {
			log.Fatalln("! Error creating nomad server node %v", i)
		} else {
			fmt.Printf("	+ nomad node %v successfully created\n", newDroplet.Name)
		}
	}

	fmt.Printf("> Start creating %v client nodes in cluster %v\n ", clientSize, name)

	for k := 0; k < clientSize; k++ {
		newDroplet, err := nodeCreationProcess(name, clusterID, encryptKey, NodeKindClient, NodeResClient, clientSize, k, region, sshKey)

		if err != nil {
			log.Fatalln("! Error creating client node %v", k)
		} else {
			fmt.Printf("	+ Node %v successfully created.\n ", newDroplet.Name)
		}
	}

	fmt.Println("> Cluster Created.")
}

// generateClusterId generate the unique cluster ID
func generateUniqueClusterID() string {
	return uuid.New().String()
}

/*
nodeCreationProcess create a node in a cluster
NodeKind can be of 2 value : [standalone, cluster]
NodeRes can be of 2 value : [server, client]
*/
func nodeCreationProcess(clusterName string, clusterID string, encryptKey string, nodeKind NodeKind, nodeRes NodeRes, nodeSize int, iteration int, region string, sshKey godo.DropletCreateSSHKey) (*godo.Droplet, error) {
	nodeName := label.GenerateClientLabel(label.Label(clusterName), iteration)

	nodeTags := getNodeTags(clusterName, clusterID, getNodeKind(nodeKind), getNodeRes(nodeRes))

	// send cluster servers ID to consul to allow him to make the cluster network
	clusterServerIDTag := tag.GetPropTag("server-id", tag.TagProp(clusterID))
	userData := ""
	if string(nodeKind) == "consul" {
		userData = consul.GetConsulServerInstallScript(clusterName, encryptKey, clusterServerIDTag, region, nodeSize)
	} else if string(nodeKind) == "nomad" {
		userData = nomad.GetNomadServerInstallScript(clusterName, encryptKey, clusterServerIDTag, region, nodeSize)
	} else {
		userData = client.GetClientInstallScript(clusterName, encryptKey, clusterServerIDTag, region)
	}

	// "#!/bin/bash \n  mkdir -p /etc/vultron;\n  cat << 'Vultron ecosystem' > /etc/vultron/README.txt"

	return CreateNode(nodeName+"-"+string(nodeKind)+"-"+string(nodeRes)+strconv.Itoa(iteration), region, sshKey, userData, nodeTags)
}

// generateClusterEncryptKey generate a 16 bytes base64 encoded key asked by Consul Gossip Encryption to secure cluster agent communications
func generateClusterEncryptKey() string {
	key := make([]byte, 64)

	// generate 16 bytes key
	_, err := rand.Read(key)
	if err != nil {
		fmt.Printf("! Error generating encrypt key: %v \n", err)
		os.Exit(1)
	}

	// base64 encoding of the key
	sEnc := base64.StdEncoding.EncodeToString([]byte(key))

	return sEnc
}

func getNodeKind(nodeKind NodeKind) NodeKind {
	if nodeKind == "consul" {
		return NodeKind(tag.GetConsulKindTag())
	} else if nodeKind == "nomad" {
		return NodeKind(tag.GetNomadKindTag())
	}

	return NodeKind(tag.GetClientKindTag())
}

func getNodeRes(nodeRes NodeRes) NodeRes {
	if nodeRes == "server" {
		return NodeRes(tag.GetServerResourceTag())
	}

	return NodeRes(tag.GetClientResourceTag())
}

func getNodeTags(clusterName string, clusterID string, nodeKind NodeKind, nodeRes NodeRes) []string {

	tags := []string{
		string(nodeKind),
		string(nodeRes),
		tag.GetPropTag("cluster-id", tag.TagProp(clusterID)),
		tag.GetPropTag("cluster-name", tag.TagProp(clusterName)),
	}

	if string(nodeRes) == tag.GetServerResourceTag() {
		tags = append(tags, tag.GetPropTag("server-id", tag.TagProp(clusterID)))
	} else {
		tags = append(tags, tag.GetPropTag("client-id", tag.TagProp(clusterID)))
	}

	return tags
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
