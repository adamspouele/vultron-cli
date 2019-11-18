package nomad

// GetNomadClientInstallScript return the nomad installation script
func GetNomadClientInstallScript(datacenterName string, encryptKey string, clusterConsulResTag string, region string) string {
	return GetNomadBaseInstallScript(datacenterName, encryptKey, clusterConsulResTag, region) + `

cat << EOF > /etc/nomad.d/client.hcl
# Enable the client
client {
	enabled = true
}
EOF

echo "Install Docker Engine"

sudo apt-get -y install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
	software-properties-common
	
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

sudo apt-key fingerprint 0EBFCD88

sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"

sudo apt-get update

sudo apt-get -y install docker-ce docker-ce-cli containerd.io

echo "Start Nomad client service"

sudo systemctl enable nomad
sudo systemctl start nomad
sudo systemctl status nomad

`
}
