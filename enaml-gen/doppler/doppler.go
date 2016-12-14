package doppler 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Doppler struct {

	/*Etcd - Descr: PEM-encoded client key Default: 
*/
	Etcd *DopplerEtcd `yaml:"etcd,omitempty"`

	/*WebsocketWriteTimeoutSeconds - Descr: Interval before a websocket write is aborted if it does not succeed Default: 60
*/
	WebsocketWriteTimeoutSeconds interface{} `yaml:"websocket_write_timeout_seconds,omitempty"`

	/*Debug - Descr: boolean value to turn on verbose logging for doppler system (dea agent & doppler server) Default: false
*/
	Debug interface{} `yaml:"debug,omitempty"`

	/*SinkIoTimeoutSeconds - Descr: I/O Timeout on sinks Default: 0
*/
	SinkIoTimeoutSeconds interface{} `yaml:"sink_io_timeout_seconds,omitempty"`

	/*LockedMemoryLimit - Descr: Size (KB) of shell's locked memory limit. Set to 'kernel' to use the kernel's default. Non-numeric values other than 'kernel', 'soft', 'hard', and 'unlimited' will result in an error. Default: unlimited
*/
	LockedMemoryLimit interface{} `yaml:"locked_memory_limit,omitempty"`

	/*MessageDrainBufferSize - Descr: Size of the internal buffer used by doppler to store messages for output to firehose or 'cf logs'. If the buffer gets full doppler will drop the messages. Default: 10000
*/
	MessageDrainBufferSize interface{} `yaml:"message_drain_buffer_size,omitempty"`

	/*BlacklistedSyslogRanges - Descr: Blacklist for IPs that should not be used as syslog drains, e.g. internal ip addresses. Default: <nil>
*/
	BlacklistedSyslogRanges interface{} `yaml:"blacklisted_syslog_ranges,omitempty"`

	/*Zone - Descr: Zone of the doppler server Default: <nil>
*/
	Zone interface{} `yaml:"zone,omitempty"`

	/*OutgoingPort - Descr: Port for outgoing log messages Default: 8081
*/
	OutgoingPort interface{} `yaml:"outgoing_port,omitempty"`

	/*MaxRetainedLogMessages - Descr: number of log messages to retain per application Default: 100
*/
	MaxRetainedLogMessages interface{} `yaml:"maxRetainedLogMessages,omitempty"`

	/*IncomingTcpPort - Descr: Port for incoming tcp messages Default: 3458
*/
	IncomingTcpPort interface{} `yaml:"incoming_tcp_port,omitempty"`

	/*SinkInactivityTimeoutSeconds - Descr: Interval before removing a sink due to inactivity Default: 3600
*/
	SinkInactivityTimeoutSeconds interface{} `yaml:"sink_inactivity_timeout_seconds,omitempty"`

	/*UnmarshallerCount - Descr: Number of parallel unmarshallers to run within Doppler Default: 5
*/
	UnmarshallerCount interface{} `yaml:"unmarshaller_count,omitempty"`

	/*Tls - Descr: TLS server certificate Default: 
*/
	Tls *DopplerTls `yaml:"tls,omitempty"`

	/*SinkDialTimeoutSeconds - Descr: Dial timeout for sinks Default: 1
*/
	SinkDialTimeoutSeconds interface{} `yaml:"sink_dial_timeout_seconds,omitempty"`

	/*SyslogSkipCertVerify - Descr: When connecting over TLS, don't verify certificates for syslog sink Default: true
*/
	SyslogSkipCertVerify interface{} `yaml:"syslog_skip_cert_verify,omitempty"`

	/*DropsondeIncomingPort - Descr: Port for incoming udp messages Default: 3457
*/
	DropsondeIncomingPort interface{} `yaml:"dropsonde_incoming_port,omitempty"`

	/*ContainerMetricTtlSeconds - Descr: TTL (in seconds) for container usage metrics Default: 120
*/
	ContainerMetricTtlSeconds interface{} `yaml:"container_metric_ttl_seconds,omitempty"`

}