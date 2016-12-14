package syslog_drain_binder 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type SyslogDrainBinderJob struct {

	/*SystemDomain - Descr: Domain reserved for CF operator, base URL where the login, uaa, and other non-user apps listen Default: <nil>
*/
	SystemDomain interface{} `yaml:"system_domain,omitempty"`

	/*SyslogDrainBinder - Descr: PEM-encoded client key Default: 
*/
	SyslogDrainBinder *SyslogDrainBinder `yaml:"syslog_drain_binder,omitempty"`

	/*Loggregator - Descr: IPs pointing to the ETCD cluster Default: <nil>
*/
	Loggregator *Loggregator `yaml:"loggregator,omitempty"`

	/*Cc - Descr: API URI of cloud controller Default: <nil>
*/
	Cc *Cc `yaml:"cc,omitempty"`

	/*Ssl - Descr: When connecting over https, ignore bad ssl certificates Default: false
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

	/*MetronEndpoint - Descr: The port used to emit dropsonde messages to the Metron agent Default: 3457
*/
	MetronEndpoint *MetronEndpoint `yaml:"metron_endpoint,omitempty"`

}