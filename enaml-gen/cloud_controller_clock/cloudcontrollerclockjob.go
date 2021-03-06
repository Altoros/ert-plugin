package cloud_controller_clock 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type CloudControllerClockJob struct {

	/*SystemDomainOrganization - Descr: The User Org that owns the system_domain, required if system_domain is defined Default: 
*/
	SystemDomainOrganization interface{} `yaml:"system_domain_organization,omitempty"`

	/*LoggerEndpoint - Descr: Whether to use ssl for logger endpoint listed at /v2/info Default: true
*/
	LoggerEndpoint *LoggerEndpoint `yaml:"logger_endpoint,omitempty"`

	/*Domain - Descr: Deprecated in favor of system_domain. Domain where cloud_controller will listen (api.domain) Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

	/*Nats - Descr: Password for cc client to connect to NATS Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

	/*SystemDomain - Descr: Domain reserved for CF operator, base URL where the login, uaa, and other non-user apps listen Default: <nil>
*/
	SystemDomain interface{} `yaml:"system_domain,omitempty"`

	/*Hm9000 - Descr: Port of the hm9000 Api Server Default: <nil>
*/
	Hm9000 *Hm9000 `yaml:"hm9000,omitempty"`

	/*Build - Descr: 'build' attribute in the /v2/info endpoint Default: 
*/
	Build interface{} `yaml:"build,omitempty"`

	/*Uaa - Descr: Used for generating SSO clients for service brokers. Default: <nil>
*/
	Uaa *Uaa `yaml:"uaa,omitempty"`

	/*Name - Descr: 'name' attribute in the /v2/info endpoint Default: 
*/
	Name interface{} `yaml:"name,omitempty"`

	/*Description - Descr: 'description' attribute in the /v2/info endpoint Default: 
*/
	Description interface{} `yaml:"description,omitempty"`

	/*DeaNext - Descr: Memory limit in mb for staging tasks Default: 1024
*/
	DeaNext *DeaNext `yaml:"dea_next,omitempty"`

	/*Login - Descr: URL of the login server Default: <nil>
*/
	Login *Login `yaml:"login,omitempty"`

	/*MetronEndpoint - Descr: The host used to emit messages to the Metron agent Default: 127.0.0.1
*/
	MetronEndpoint *MetronEndpoint `yaml:"metron_endpoint,omitempty"`

	/*Version - Descr: 'version' attribute in the /v2/info endpoint Default: 0
*/
	Version interface{} `yaml:"version,omitempty"`

	/*Ccdb - Descr: Maximum connections for Sequel Default: 25
*/
	Ccdb *Ccdb `yaml:"ccdb,omitempty"`

	/*RequestTimeoutInSeconds - Descr: Timeout for requests in seconds. Default: 900
*/
	RequestTimeoutInSeconds interface{} `yaml:"request_timeout_in_seconds,omitempty"`

	/*Ssl - Descr: specifies that the job is allowed to skip ssl cert verification Default: false
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

	/*AppDomains - Descr: Array of domains for user apps (example: 'user.app.space.foo', a user app called 'neat' will listen at 'http://neat.user.app.space.foo') Default: <nil>
*/
	AppDomains interface{} `yaml:"app_domains,omitempty"`

	/*Cc - Descr: The file descriptors made available to each app instance Default: 16384
*/
	Cc *Cc `yaml:"cc,omitempty"`

	/*SupportAddress - Descr: 'support' attribute in the /v2/info endpoint Default: 
*/
	SupportAddress interface{} `yaml:"support_address,omitempty"`

}