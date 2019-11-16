#!/bin/bash

# reference : https://learn.hashicorp.com/consul/datacenter-deploy/deployment-guide

apt-get update

apt-get install -y unzip

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
datacenter = "meshectares-datacenter-01"
data_dir = "/opt/consul"
encrypt = "4m40FhL9yOIdP4SqYOvQVGPidmWgnOe2pv4Yjk2H/KA="
{
    "retry_join": ["provider=digitalocean region=ams3 tag_name=consul:server api_token=a672e9dc039db44186f3b9e1bd4a2ac0f4c44844a037b22d2e51770cc5164dc7"]
}

performance {
  raft_multiplier = 1
}
EOF

echo "Following steps will set current consul agent to server mode."

echo "Create a configuration file at /etc/consul.d/server.hcl"

sudo mkdir --parents /etc/consul.d
sudo touch /etc/consul.d/server.hcl
sudo chown --recursive consul:consul /etc/consul.d
sudo chmod 640 /etc/consul.d/server.hcl

echo "configure server.hcl"

cat <<EOF >/etc/consul.d/server.hcl
server = true
bootstrap_expect = 3
EOF

echo "Following steps will set current consul agent to client mode."

echo "Create a configuration file at /etc/consul.d/server.hcl"

echo "Check Consul service"

sudo systemctl enable consul
sudo systemctl start consul
sudo systemctl status consul

