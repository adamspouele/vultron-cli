package client

import nomad "github.com/adamspouele/vultron-cli/package/nomad/script"

// GetClientInstallScript return the client installation script
func GetClientInstallScript(datacenterName string, encryptKey string, clusterServersTag string, region string) string {
	return nomad.GetNomadClientInstallScript(datacenterName, encryptKey, clusterServersTag, region) + `

echo "Check Consul service"

sudo systemctl enable consul
sudo systemctl start consul
sudo systemctl status consul

`
}
