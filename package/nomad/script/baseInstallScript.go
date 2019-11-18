package nomad

import consul "github.com/adamspouele/vultron-cli/package/consul/script"

// GetNomadBaseInstallScript return the nomad installation script
func GetNomadBaseInstallScript(datacenterName string, encryptKey string, clusterConsulResTag string, region string) string {
	return consul.GetConsulClientInstallScript(datacenterName, encryptKey, clusterConsulResTag, region) + `

echo "Download Nomad"
export NOMAD_VERSION="0.10.1"
curl --silent --remote-name https://releases.hashicorp.com/nomad/${NOMAD_VERSION}/nomad_${NOMAD_VERSION}_linux_amd64.zip

echo "Intall Nomad"

unzip nomad_${NOMAD_VERSION}_linux_amd64.zip
sudo chown root:root nomad
sudo mv nomad /usr/local/bin/
nomad version

nomad -autocomplete-install
complete -C /usr/local/bin/nomad nomad

sudo mkdir --parents /opt/nomad

echo "Configure systemd"

cat << EOF > /etc/systemd/system/nomad.service
[Unit]
Description=Nomad
Documentation=https://nomadproject.io/docs/
Wants=network-online.target
After=network-online.target

[Service]
ExecReload=/bin/kill -HUP $MAINPID
ExecStart=/usr/local/bin/nomad agent -config /etc/nomad.d
KillMode=process
KillSignal=SIGINT
LimitNOFILE=infinity
LimitNPROC=infinity
Restart=on-failure
RestartSec=2
StartLimitBurst=3
StartLimitIntervalSec=10
TasksMax=infinity

[Install]
WantedBy=multi-user.target
EOF

echo "Configure Nomad"

echo "configure nomad.hcl"

sudo mkdir --parents /etc/nomad.d
sudo chmod 700 /etc/nomad.d
sudo touch /etc/nomad.d/nomad.hcl

cat << EOF > /etc/nomad.d/nomad.hcl
datacenter = "` + datacenterName + `"
data_dir = "/opt/nomad"
EOF

`
}
