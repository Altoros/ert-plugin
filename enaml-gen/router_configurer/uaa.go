package router_configurer 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Uaa struct {

	/*CaCert - Descr: Certificate authority for communication between clients and uaa. Default: 
*/
	CaCert interface{} `yaml:"ca_cert,omitempty"`

	/*TlsPort - Descr: Port on which UAA is listening for TLS connections. This is required for obtaining an OAuth token for Routing API. Default: <nil>
*/
	TlsPort interface{} `yaml:"tls_port,omitempty"`

}