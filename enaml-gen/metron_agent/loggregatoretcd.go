package metron_agent 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type LoggregatorEtcd struct {

	/*CaCert - Descr: PEM-encoded CA certificate Default: 
*/
	CaCert interface{} `yaml:"ca_cert,omitempty"`

	/*RequireSsl - Descr: Enable ssl for all communication with etcd Default: false
*/
	RequireSsl interface{} `yaml:"require_ssl,omitempty"`

	/*Maxconcurrentrequests - Descr: Number of concurrent requests to ETCD Default: 10
*/
	Maxconcurrentrequests interface{} `yaml:"maxconcurrentrequests,omitempty"`

	/*Machines - Descr: IPs pointing to the ETCD cluster Default: <nil>
*/
	Machines interface{} `yaml:"machines,omitempty"`

}