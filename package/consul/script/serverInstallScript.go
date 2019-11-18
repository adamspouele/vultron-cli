package consul

import "strconv"

// GetConsulServerInstallScript return the consul installation script
func GetConsulServerInstallScript(datacenterName string, encryptKey string, clusterConsulResTag string, region string, serverSize int) string {
	return GetConsulBaseInstallScript(datacenterName, encryptKey, clusterConsulResTag, region) + `
echo "configure server.hcl"

cat <<EOF >/etc/consul.d/server.hcl
server = true
bootstrap_expect = ` + strconv.Itoa(serverSize) + `
EOF

echo "This Consul instance is on server mode."

echo "Check Consul service"

sudo systemctl enable consul
sudo systemctl start consul
sudo systemctl status consul
`
}
