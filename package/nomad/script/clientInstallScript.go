package nomad

// GetNomadClientInstallScript return the nomad installation script
func GetNomadClientInstallScript(datacenterName string, encryptKey string, clusterServersTag string, region string) string {
	return GetConsulBaseInstallScript(datacenterName, encryptKey, clusterServersTag, region) + `

echo "Start Nomad client service"

sudo systemctl enable nomad
sudo systemctl start nomad
sudo systemctl status nomad

`
}
