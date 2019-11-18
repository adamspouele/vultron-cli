package consul

// GetConsulClientInstallScript return the consul installation script
func GetConsulClientInstallScript(datacenterName string, encryptKey string, clusterConsulResTag string, region string) string {
	return GetConsulBaseInstallScript(datacenterName, encryptKey, clusterConsulResTag, region) + `

echo "Check Consul service"

sudo systemctl enable consul
sudo systemctl start consul
sudo systemctl status consul

`
}
