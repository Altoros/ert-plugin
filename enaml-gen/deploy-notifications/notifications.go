package deploy_notifications 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Notifications struct {

	/*Smtp - Descr: Username for the SMTP host Default: <nil>
*/
	Smtp *Smtp `yaml:"smtp,omitempty"`

	/*Sender - Descr: Sender of the notification Default: <nil>
*/
	Sender interface{} `yaml:"sender,omitempty"`

	/*EncryptionKey - Descr: key used to encrypt unsubscribe IDs Default: <nil>
*/
	EncryptionKey interface{} `yaml:"encryption_key,omitempty"`

	/*Network - Descr: Network used to host application Default: <nil>
*/
	Network interface{} `yaml:"network,omitempty"`

	/*Organization - Descr: Organization that hosts the app Default: <nil>
*/
	Organization interface{} `yaml:"organization,omitempty"`

	/*Space - Descr: Space that hosts the app Default: <nil>
*/
	Space interface{} `yaml:"space,omitempty"`

	/*ErrorOnMisconfiguration - Descr: Throw error on service misconfiguration during deployment Default: true
*/
	ErrorOnMisconfiguration interface{} `yaml:"error_on_misconfiguration,omitempty"`

	/*Cf - Descr: Password of the CF admin user Default: <nil>
*/
	Cf *Cf `yaml:"cf,omitempty"`

	/*InstanceCount - Descr: number of instances of service to run Default: <nil>
*/
	InstanceCount interface{} `yaml:"instance_count,omitempty"`

	/*BuildpackUrl - Descr: Optional parameter that specifies the url of the buildpack to use.  Helpful if your environment does not have the latest Go buildpack. Default: <nil>
*/
	BuildpackUrl interface{} `yaml:"buildpack_url,omitempty"`

	/*Database - Descr: URL pointing to database Default: <nil>
*/
	Database *Database `yaml:"database,omitempty"`

	/*Uaa - Descr: Client id of the UAA Default: <nil>
*/
	Uaa *Uaa `yaml:"uaa,omitempty"`

	/*SyslogUrl - Descr: External log service URL Default: <nil>
*/
	SyslogUrl interface{} `yaml:"syslog_url,omitempty"`

	/*AppDomain - Descr: Domain used to host application Default: <nil>
*/
	AppDomain interface{} `yaml:"app_domain,omitempty"`

	/*EnableDiego - Descr: Enable deployment to diego Default: <nil>
*/
	EnableDiego interface{} `yaml:"enable_diego,omitempty"`

	/*DefaultTemplate - Descr: default template to use for the service Default: <nil>
*/
	DefaultTemplate interface{} `yaml:"default_template,omitempty"`

}