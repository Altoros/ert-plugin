package proxy 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Proxy struct {

	/*HealthPort - Descr: Port for checking the health of the proxy process Default: 1936
*/
	HealthPort interface{} `yaml:"health_port,omitempty"`

	/*Port - Descr: Port for the proxy to listen on Default: 3306
*/
	Port interface{} `yaml:"port,omitempty"`

	/*ProxyIps - Descr: List of IP addresses for all proxy jobs Default: <nil>
*/
	ProxyIps interface{} `yaml:"proxy_ips,omitempty"`

	/*ApiUsername - Descr: Username for Basic Auth used to secure API Default: <nil>
*/
	ApiUsername interface{} `yaml:"api_username,omitempty"`

	/*ApiPassword - Descr: Password for Basic Auth used to secure API Default: <nil>
*/
	ApiPassword interface{} `yaml:"api_password,omitempty"`

	/*ApiForceHttps - Descr: Redirect all HTTP requests to the API to HTTPS Default: true
*/
	ApiForceHttps interface{} `yaml:"api_force_https,omitempty"`

	/*ApiPort - Descr: Port for the proxy API to listen on Default: 80
*/
	ApiPort interface{} `yaml:"api_port,omitempty"`

	/*ArbitratorIp - Descr: List of IP addresses for the arbitrator nodes of the MySQL cluster Default: no-arbitrator-ip
*/
	ArbitratorIp interface{} `yaml:"arbitrator_ip,omitempty"`

	/*HealthcheckTimeoutMillis - Descr: Timeout (milliseconds) before assuming a backend is unhealthy Default: 5000
*/
	HealthcheckTimeoutMillis interface{} `yaml:"healthcheck_timeout_millis,omitempty"`

}