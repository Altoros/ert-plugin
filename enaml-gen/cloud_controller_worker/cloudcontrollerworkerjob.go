package cloud_controller_worker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type CloudControllerWorkerJob struct {

	/*Domain - Descr: Deprecated in favor of system_domain. Domain where cloud_controller will listen (api.domain) Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

	/*Ssl - Descr: specifies that the job is allowed to skip ssl cert verification Default: false
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

	/*Nats - Descr: Username for cc client to connect to NATS Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

	/*DeaNext - Descr: Disk limit in mb for staging tasks Default: 6144
*/
	DeaNext *DeaNext `yaml:"dea_next,omitempty"`

	/*Build - Descr: 'build' attribute in the /v2/info endpoint Default: 
*/
	Build interface{} `yaml:"build,omitempty"`

	/*Login - Descr: whether use login as the authorization endpoint or not Default: true
*/
	Login *Login `yaml:"login,omitempty"`

	/*Uaa - Descr: Used to grant scope for SSO clients for service brokers Default: openid,cloud_controller_service_permissions.read
*/
	Uaa *Uaa `yaml:"uaa,omitempty"`

	/*AppDomains - Descr: Array of domains for user apps (example: 'user.app.space.foo', a user app called 'neat' will listen at 'http://neat.user.app.space.foo') Default: <nil>
*/
	AppDomains interface{} `yaml:"app_domains,omitempty"`

	/*SystemDomain - Descr: Domain reserved for CF operator, base URL where the login, uaa, and other non-user apps listen Default: <nil>
*/
	SystemDomain interface{} `yaml:"system_domain,omitempty"`

	/*Cc - Descr: The environment name used by NewRelic Default: development
*/
	Cc *Cc `yaml:"cc,omitempty"`

	/*LoggerEndpoint - Descr: Port for logger endpoint listed at /v2/info Default: 443
*/
	LoggerEndpoint *LoggerEndpoint `yaml:"logger_endpoint,omitempty"`

	/*SupportAddress - Descr: 'support' attribute in the /v2/info endpoint Default: 
*/
	SupportAddress interface{} `yaml:"support_address,omitempty"`

	/*Name - Descr: 'name' attribute in the /v2/info endpoint Default: 
*/
	Name interface{} `yaml:"name,omitempty"`

	/*RequestTimeoutInSeconds - Descr: Timeout for requests in seconds. Default: 900
*/
	RequestTimeoutInSeconds interface{} `yaml:"request_timeout_in_seconds,omitempty"`

	/*SystemDomainOrganization - Descr: The User Org that owns the system_domain, required if system_domain is defined Default: 
*/
	SystemDomainOrganization interface{} `yaml:"system_domain_organization,omitempty"`

	/*NfsServer - Descr: The location at which to mount the nfs share Default: /var/vcap/nfs
*/
	NfsServer *NfsServer `yaml:"nfs_server,omitempty"`

	/*Version - Descr: 'version' attribute in the /v2/info endpoint Default: 0
*/
	Version interface{} `yaml:"version,omitempty"`

	/*Hm9000 - Descr: Port of the hm9000 Api Server Default: <nil>
*/
	Hm9000 *Hm9000 `yaml:"hm9000,omitempty"`

	/*Ccdb - Descr: Contains the name of the database on the database server Default: <nil>
*/
	Ccdb *Ccdb `yaml:"ccdb,omitempty"`

	/*Description - Descr: 'description' attribute in the /v2/info endpoint Default: 
*/
	Description interface{} `yaml:"description,omitempty"`

	/*MetronEndpoint - Descr: The host used to emit messages to the Metron agent Default: 127.0.0.1
*/
	MetronEndpoint *MetronEndpoint `yaml:"metron_endpoint,omitempty"`

}