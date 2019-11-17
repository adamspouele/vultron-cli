package client

// GetBaseInstallScript common base installation script of all nodes of the cluster
func GetBaseInstallScript() string {
	return `#!/bin/bash
apt-get update

apt-get install -y unzip

mkdir /etc/vultron

set -x
exec > >(tee /var/log/user-data.log|logger -t user-data ) 2>&1
echo BEGIN
date '+%Y-%m-%d %H:%M:%S'

`
}
