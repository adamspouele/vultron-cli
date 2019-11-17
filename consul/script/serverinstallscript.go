package consul

// GetConsulServerInstallScript return the consul installation script
func GetConsulServerInstallScript() string {
	return GetConsulBaseInstallScript() + `
	echo "configure server.hcl"

	cat <<EOF >/etc/consul.d/server.hcl
	server = true
	bootstrap_expect = 3
	EOF
	
	echo "Following steps will set current consul agent to client mode."

	echo "Check Consul service"

	sudo systemctl enable consul
	sudo systemctl start consul
	sudo systemctl status consul

	`
}
