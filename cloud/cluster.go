package cloud

import (
	"fmt"
	"github.com/digitalocean/godo"
	"strconv"
)

// CreateCluster create a cluster in the cloud
func CreateCluster(name string, region string, serverSize int, clientSize int, sshKey godo.DropletCreateSSHKey) {

	fmt.Printf("> Plan deployment of %v server nodes \n", serverSize)

	fmt.Printf("> Start creating %v server nodes in %v cluster \n", clientSize, name)

	for i := 0; i < serverSize; i++ {
		nodeTags := []string{"vultron:test"}
		newDroplet, err := CreateNode(name+"-"+strconv.Itoa(i), region, sshKey, "", nodeTags)

		if err != nil {
			fmt.Printf("! Error creating server node %v \n", i)
		} else {
			fmt.Printf("> Node %v successfully created. \n", newDroplet.Name)
		}
	}

	fmt.Printf("> Start creating %v client nodes in cluster %v \n", clientSize, name)

	for k := 0; k < clientSize; k++ {

	}

	fmt.Printf("> Cluster Created. \n")
}

// ListClusterNodes list cluster nodes
func ListClusterNodes(clusterTag string) ([]godo.Droplet, error) {
	client, ctx, err := GetDoClient()
	opt := &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	}

	if err != nil {
		fmt.Printf("error authenticating")
	}

	droplets, _, err := client.Droplets.ListByTag(ctx, clusterTag, opt)

	return droplets, err
}

// ExplainCluster explain a cluster
func ExplainCluster(clusterTag string) {

	droplets, err := ListClusterNodes(clusterTag)

	if err != nil {
		fmt.Printf("! Error while explaining cluster %v \n", clusterTag)
	}

	for i := 0; i < len(droplets); i++ {
		fmt.Printf("# %v : ID='%v' status='%v' name='%v' region='%v' memory='%v' vcpus='%v' image='%v' disk='%v' size='%v' created='%v' \n",
			i, droplets[i].ID, droplets[i].Status, droplets[i].Name, droplets[i].Region.Slug, droplets[i].Memory,
			droplets[i].Vcpus, droplets[i].Image.Slug, droplets[i].Disk, droplets[i].Size.Slug, droplets[i].Created)
	}
}

// DeleteCluster delete a cluster by removing all cluster's nodes
func DeleteCluster(clusterTag string) (*godo.Response, error) {

	fmt.Printf("> Delete cluster %v \n", clusterTag)

	client, ctx, _ := GetDoClient()

	response, err := client.Droplets.DeleteByTag(ctx, clusterTag)

	fmt.Printf("%v", response.Response.Status)

	if err != nil {
		fmt.Printf("! Error deleting cluster [%v] : %v \n", clusterTag, err)
	}

	fmt.Printf("Done. \n")

	return response, err
}
