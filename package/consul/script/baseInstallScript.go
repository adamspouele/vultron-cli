package consul

import client "github.com/adamspouele/vultron-cli/package/script"

// GetConsulBaseInstallScript return the consul installation script
func GetConsulBaseInstallScript(datacenterName string, encryptKey string, clusterConsulResTag string, region string) string {
	return client.GetBaseInstallScript() + `

# reference : https://learn.hashicorp.com/consul/datacenter-deploy/deployment-guide

CONSUL_VERSION="1.6.1"

curl --silent --remote-name https://releases.hashicorp.com/consul/${CONSUL_VERSION}/consul_${CONSUL_VERSION}_linux_amd64.zip

echo "Install Consul"

unzip consul_${CONSUL_VERSION}_linux_amd64.zip


sudo chown root:root consul

sudo mv consul /usr/local/bin/

consul --version

consul -autocomplete-install

complete -C /usr/local/bin/consul consul

sudo useradd --system --home /etc/consul.d --shell /bin/false consul

sudo mkdir --parents /opt/consul

sudo chown --recursive consul:consul /opt/consul

echo "Configure systemd"

cat <<EOF >/etc/systemd/system/consul.service
[Unit]
Description="HashiCorp Consul - A service mesh solution"
Documentation=https://www.consul.io/
Requires=network-online.target
After=network-online.target
ConditionFileNotEmpty=/etc/consul.d/consul.hcl

[Service]
Type=notify
User=consul
Group=consul
ExecStart=/usr/local/bin/consul agent -config-dir=/etc/consul.d/
ExecReload=/usr/local/bin/consul reload
KillMode=process
Restart=on-failure
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
EOF

echo "Create a configuration file at /etc/consul.d/consul.hcl"

mkdir --parents /etc/consul.d
touch /etc/consul.d/consul.hcl
chown --recursive consul:consul /etc/consul.d
chmod 640 /etc/consul.d/consul.hcl

echo "configure consul.hcl"

cat <<EOF >/etc/consul.d/consul.hcl
datacenter = "` + datacenterName + `"
data_dir = "/opt/consul"
log_level = "DEBUG"
enable_syslog = true
encrypt = "` + encryptKey + `"
bind_addr = "$(curl http://169.254.169.254/metadata/v1/interfaces/private/0/ipv4/address)"
retry_join = ["provider=digitalocean region=` + region + ` tag_name=` + clusterConsulResTag + ` api_token=a672e9dc039db44186f3b9e1bd4a2ac0f4c44844a037b22d2e51770cc5164dc7"]

performance {
  raft_multiplier = 1
}
EOF

`
}
