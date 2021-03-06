package routing_api 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type RoutingApiJob struct {

	/*RoutingApi - Descr: Password used for connecting to SQL database Default: <nil>
*/
	RoutingApi *RoutingApi `yaml:"routing_api,omitempty"`

	/*SkipSslValidation - Descr: Skip TLS verification when talking to UAA Default: false
*/
	SkipSslValidation interface{} `yaml:"skip_ssl_validation,omitempty"`

	/*DnsHealthCheckHost - Descr: Host to ping for confirmation of DNS resolution Default: consul.service.cf.internal
*/
	DnsHealthCheckHost interface{} `yaml:"dns_health_check_host,omitempty"`

	/*Uaa - Descr: Port on which UAA is listening for TLS connections. This is required for obtaining a key to verify client OAuth tokens. Default: <nil>
*/
	Uaa *Uaa `yaml:"uaa,omitempty"`

	/*Consul - Descr: comma-separated list of consul server URLs (scheme://ip:port) Default: http://127.0.0.1:8500
*/
	Consul *Consul `yaml:"consul,omitempty"`

	/*Metron - Descr: The port used to emit dropsonde messages to the Metron agent. Default: 3457
*/
	Metron *Metron `yaml:"metron,omitempty"`

}