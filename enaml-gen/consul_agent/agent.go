package consul_agent 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Agent struct {

	/*ProtocolVersion - Descr: The Consul protocol to use. Default: 2
*/
	ProtocolVersion interface{} `yaml:"protocol_version,omitempty"`

	/*Mode - Descr: Mode to run the agent in. (client or server) Default: client
*/
	Mode interface{} `yaml:"mode,omitempty"`

	/*Servers - Descr: WAN server addresses to join. Default: []
*/
	Servers *Servers `yaml:"servers,omitempty"`

	/*Datacenter - Descr: Name of the agent's datacenter. Default: dc1
*/
	Datacenter interface{} `yaml:"datacenter,omitempty"`

	/*DnsConfig - Descr: When allow_stale is specified, this is used to limit how stale results are allowed to be. Default: 5s
*/
	DnsConfig *DnsConfig `yaml:"dns_config,omitempty"`

	/*Domain - Descr: Domain suffix for DNS Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

	/*LogLevel - Descr: Agent log level. Default: info
*/
	LogLevel interface{} `yaml:"log_level,omitempty"`

	/*Services - Descr: Map of consul service definitions. Default: map[]
*/
	Services interface{} `yaml:"services,omitempty"`

}