package etcd_testconsumer 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Etcd struct {

	/*CaCert - Descr: PEM-encoded CA certificate Default: 
*/
	CaCert interface{} `yaml:"ca_cert,omitempty"`

	/*ClientCert - Descr: PEM-encoded client certificate Default: 
*/
	ClientCert interface{} `yaml:"client_cert,omitempty"`

	/*ClientKey - Descr: PEM-encoded client key Default: 
*/
	ClientKey interface{} `yaml:"client_key,omitempty"`

	/*DnsHealthCheckHost - Descr: Host to ping for confirmation of DNS resolution Default: consul.service.cf.internal
*/
	DnsHealthCheckHost interface{} `yaml:"dns_health_check_host,omitempty"`

	/*RequireSsl - Descr: enable ssl for all communication with etcd Default: false
*/
	RequireSsl interface{} `yaml:"require_ssl,omitempty"`

	/*Machines - Descr: Addresses of etcd machines Default: <nil>
*/
	Machines interface{} `yaml:"machines,omitempty"`

}