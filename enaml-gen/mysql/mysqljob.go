package mysql 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type MysqlJob struct {

	/*CfMysql - Descr: Enable binlogs across all nodes Default: true
*/
	CfMysql *CfMysql `yaml:"cf_mysql,omitempty"`

	/*SyslogAggregator - Descr: IP address for syslog aggregator Default: <nil>
*/
	SyslogAggregator *SyslogAggregator `yaml:"syslog_aggregator,omitempty"`

}