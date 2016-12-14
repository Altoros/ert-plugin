package arbitrator 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Mysql struct {

	/*Port - Descr: Port the mysql server should bind to Default: 3306
*/
	Port interface{} `yaml:"port,omitempty"`

	/*GaleraPort - Descr: Port which garbd listens on Default: 4567
*/
	GaleraPort interface{} `yaml:"galera_port,omitempty"`

	/*ClusterIps - Descr: List of nodes.  Must have the same number of ips as there are nodes in the cluster Default: <nil>
*/
	ClusterIps interface{} `yaml:"cluster_ips,omitempty"`

	/*GaleraHealthcheck - Descr: Username used by the sidecar endpoints for Basic Auth Default: <nil>
*/
	GaleraHealthcheck *GaleraHealthcheck `yaml:"galera_healthcheck,omitempty"`

	/*AdminUsername - Descr: Username for the MySQL server admin user Default: root
*/
	AdminUsername interface{} `yaml:"admin_username,omitempty"`

	/*AdminPassword - Descr: Password for the MySQL server admin user Default: <nil>
*/
	AdminPassword interface{} `yaml:"admin_password,omitempty"`

}