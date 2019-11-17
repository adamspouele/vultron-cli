package consul

// GetConsulClientInstallScript return the consul installation script
func GetConsulClientInstallScript(datacenterName string, encryptKey string, clusterServersTag string, region string) string {
	return GetConsulBaseInstallScript(datacenterName, encryptKey, clusterServersTag, region) + `

echo "Check Consul service"

sudo systemctl enable consul
sudo systemctl start consul
sudo systemctl status consul

`
}
