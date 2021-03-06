package cf_mysql_broker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Broker struct {

	/*AuthPassword - Descr: Broker's basic auth password Default: <nil>
*/
	AuthPassword interface{} `yaml:"auth_password,omitempty"`

	/*CookieSecret - Descr: A unique secret key, used to sign sessions Default: <nil>
*/
	CookieSecret interface{} `yaml:"cookie_secret,omitempty"`

	/*Services - Descr: Services and plans offered by the broker Default: <nil>
*/
	Services interface{} `yaml:"services,omitempty"`

	/*QuotaEnforcer - Descr: In seconds, the interval that the Quota Enforcer pauses between checks for violators and reformers Default: 1
*/
	QuotaEnforcer *QuotaEnforcer `yaml:"quota_enforcer,omitempty"`

	/*AuthUsername - Descr: Broker's basic auth username Default: <nil>
*/
	AuthUsername interface{} `yaml:"auth_username,omitempty"`

	/*SslEnabled - Descr: Determines use of https in dashboard url and in callback uri for calls to UAA Default: true
*/
	SslEnabled interface{} `yaml:"ssl_enabled,omitempty"`

	/*MaxUserConnectionsDefault - Descr: number of user connections to allow in a plan if not specified Default: 40
*/
	MaxUserConnectionsDefault interface{} `yaml:"max_user_connections_default,omitempty"`

}