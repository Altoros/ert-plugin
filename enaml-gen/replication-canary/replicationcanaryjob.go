package replication_canary 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type ReplicationCanaryJob struct {

	/*CfMysql - Descr: Domain of the route registered for the UI via NATS (with the router in cf-release) Default: <nil>
*/
	CfMysql *CfMysql `yaml:"cf_mysql,omitempty"`

	/*MysqlMonitoring - Descr: Username of the UAA client used to create the notifications client Default: admin
*/
	MysqlMonitoring *MysqlMonitoring `yaml:"mysql-monitoring,omitempty"`

	/*SyslogAggregator - Descr: IP address for syslog aggregator Default: <nil>
*/
	SyslogAggregator *SyslogAggregator `yaml:"syslog_aggregator,omitempty"`

	/*Domain - Descr: Domain reserved for CF operator, base URL where the login, uaa, and other non-user apps listen Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

}