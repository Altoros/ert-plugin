package router_configurer 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type RouterConfigurerJob struct {

	/*RouterConfigurer - Descr: Log level Default: info
*/
	RouterConfigurer *RouterConfigurer `yaml:"router_configurer,omitempty"`

	/*SkipSslValidation - Descr: Skip TLS verification when talking to UAA Default: false
*/
	SkipSslValidation interface{} `yaml:"skip_ssl_validation,omitempty"`

	/*Metron - Descr: The port used to emit dropsonde messages to the Metron agent. Default: 3457
*/
	Metron *Metron `yaml:"metron,omitempty"`

	/*Uaa - Descr: Port on which UAA is listening for TLS connections. This is required for obtaining an OAuth token for Routing API. Default: <nil>
*/
	Uaa *Uaa `yaml:"uaa,omitempty"`

	/*RoutingApi - Descr: When false, Routing API requires OAuth tokens for authentication. Default: false
*/
	RoutingApi *RoutingApi `yaml:"routing_api,omitempty"`

	/*DnsHealthCheckHost - Descr: Host to ping for confirmation of DNS resolution Default: consul.service.cf.internal
*/
	DnsHealthCheckHost interface{} `yaml:"dns_health_check_host,omitempty"`

}