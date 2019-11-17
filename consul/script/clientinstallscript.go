package consul

// GetConsulClientInstallScript return the consul installation script
func GetConsulClientInstallScript() string {
	return GetConsulBaseInstallScript() + `

echo "Check Consul service"

sudo systemctl enable consul
sudo systemctl start consul
sudo systemctl status consul

`
}
