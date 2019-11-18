package nomad

import "strconv"

// GetConsulServerInstallScript return the consul installation script
func GetNomadServerInstallScript(datacenterName string, encryptKey string, clusterConsulResTag string, region string, nomadServerSize int) string {
	return GetNomadBaseInstallScript(datacenterName, encryptKey, clusterConsulResTag, region) + `

echo "Server configuration"

echo "configure server.hcl"

cat << EOF > /etc/nomad.d/server.hcl
# Increase log verbosity
log_level = "DEBUG"

# Enable the server
server {
	enabled = true
	bootstrap_expect = ` + strconv.Itoa(nomadServerSize) + `
}
EOF

echo "Start Nomad server service"

sudo systemctl enable nomad
sudo systemctl start nomad
sudo systemctl status nomad
`
}
