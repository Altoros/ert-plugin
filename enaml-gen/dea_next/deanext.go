package dea_next 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type DeaNext struct {

	/*Stacks - Descr: An array of stacks, specifying the name and package path. Default: [map[name:cflinuxfs2 package_path:/var/vcap/packages/rootfs_cflinuxfs2/rootfs]]
*/
	Stacks interface{} `yaml:"stacks,omitempty"`

	/*Zone - Descr: The Availability Zone Default: default
*/
	Zone interface{} `yaml:"zone,omitempty"`

	/*InstanceMinCpuShareLimit - Descr: The minimum number of CPU shares that can be given to an app Default: 1
*/
	InstanceMinCpuShareLimit interface{} `yaml:"instance_min_cpu_share_limit,omitempty"`

	/*EnableSsl - Descr: enable ssl for all communication with DEAs Default: true
*/
	EnableSsl interface{} `yaml:"enable_ssl,omitempty"`

	/*StreamingTimeout - Descr:  Default: 60
*/
	StreamingTimeout interface{} `yaml:"streaming_timeout,omitempty"`

	/*DirectoryServerProtocol - Descr: The protocol to use when communicating with the directory server ("http" or "https") Default: https
*/
	DirectoryServerProtocol interface{} `yaml:"directory_server_protocol,omitempty"`

	/*AllowNetworks - Descr:  Default: <nil>
*/
	AllowNetworks interface{} `yaml:"allow_networks,omitempty"`

	/*InstanceBandwidthLimit - Descr: Network bandwidth burst limit for running instances in bytes Default: <nil>
*/
	InstanceBandwidthLimit *InstanceBandwidthLimit `yaml:"instance_bandwidth_limit,omitempty"`

	/*InstanceMaxCpuShareLimit - Descr: The maximum number of CPU shares that can be given to an app Default: 256
*/
	InstanceMaxCpuShareLimit interface{} `yaml:"instance_max_cpu_share_limit,omitempty"`

	/*InstanceMemoryToCpuShareRatio - Descr: Controls the relationship between app memory and cpu shares. app_cpu_shares = app_memory / cpu_share_factor Default: 8
*/
	InstanceMemoryToCpuShareRatio interface{} `yaml:"instance_memory_to_cpu_share_ratio,omitempty"`

	/*DiskOvercommitFactor - Descr:  Default: 1
*/
	DiskOvercommitFactor interface{} `yaml:"disk_overcommit_factor,omitempty"`

	/*InstanceNprocLimit - Descr: Limit on nproc for an instance container Default: 512
*/
	InstanceNprocLimit interface{} `yaml:"instance_nproc_limit,omitempty"`

	/*Mtu - Descr: Interface MTU size Default: 1500
*/
	Mtu interface{} `yaml:"mtu,omitempty"`

	/*DiskMb - Descr:  Default: 32000
*/
	DiskMb interface{} `yaml:"disk_mb,omitempty"`

	/*MemoryMb - Descr:  Default: 8000
*/
	MemoryMb interface{} `yaml:"memory_mb,omitempty"`

	/*LoggingLevel - Descr: Log level for DEA. Default: debug
*/
	LoggingLevel interface{} `yaml:"logging_level,omitempty"`

	/*StagingDiskInodeLimit - Descr: Limit on inodes for a staging container Default: 200000
*/
	StagingDiskInodeLimit interface{} `yaml:"staging_disk_inode_limit,omitempty"`

	/*KernelNetworkTuningEnabled - Descr: with latest kernel version, no kernel network tunings allowed with in warden cpi containers Default: true
*/
	KernelNetworkTuningEnabled interface{} `yaml:"kernel_network_tuning_enabled,omitempty"`

	/*EvacuationBailOutTimeInSeconds - Descr: Duration to wait before shutting down, in seconds. Default: 115
*/
	EvacuationBailOutTimeInSeconds interface{} `yaml:"evacuation_bail_out_time_in_seconds,omitempty"`

	/*StagingDiskLimitMb - Descr: Disk limit in mb for staging tasks Default: 6144
*/
	StagingDiskLimitMb interface{} `yaml:"staging_disk_limit_mb,omitempty"`

	/*DefaultHealthCheckTimeout - Descr: Default timeout for application to start Default: 60
*/
	DefaultHealthCheckTimeout interface{} `yaml:"default_health_check_timeout,omitempty"`

	/*SslPort - Descr: SSL port for DEA Default: 22443
*/
	SslPort interface{} `yaml:"ssl_port,omitempty"`

	/*AllowHostAccess - Descr: Allows warden containers to access the DEA host via its IP Default: false
*/
	AllowHostAccess interface{} `yaml:"allow_host_access,omitempty"`

	/*HeartbeatIntervalInSeconds - Descr: frequency of heartbeats in seconds. Default: 10
*/
	HeartbeatIntervalInSeconds interface{} `yaml:"heartbeat_interval_in_seconds,omitempty"`

	/*PostSetupHook - Descr: DEPRECATED: a single line of bash to prepend to the start command Default: <nil>
*/
	PostSetupHook interface{} `yaml:"post_setup_hook,omitempty"`

	/*StagingMemoryLimitMb - Descr: Memory limit in mb for staging tasks Default: 1024
*/
	StagingMemoryLimitMb interface{} `yaml:"staging_memory_limit_mb,omitempty"`

	/*RlimitCore - Descr: Maximum size of core file in bytes. 0 represents no core dump files can be created, and -1 represents no size limits. Default: 0
*/
	RlimitCore interface{} `yaml:"rlimit_core,omitempty"`

	/*ServerCert - Descr: PEM-encoded server certificate Default: <nil>
*/
	ServerCert interface{} `yaml:"server_cert,omitempty"`

	/*AdvertiseIntervalInSeconds - Descr: frequency of staging & DEA advertisments in seconds. Default: 5
*/
	AdvertiseIntervalInSeconds interface{} `yaml:"advertise_interval_in_seconds,omitempty"`

	/*DnsServers - Descr: List of nameservers to use in containers Default: <nil>
*/
	DnsServers interface{} `yaml:"dns_servers,omitempty"`

	/*StagingBandwidthLimit - Descr: Network bandwidth burst limit for staging tasks in bytes Default: <nil>
*/
	StagingBandwidthLimit *StagingBandwidthLimit `yaml:"staging_bandwidth_limit,omitempty"`

	/*CrashLifetimeSecs - Descr: Crashed app lifetime in seconds Default: 3600
*/
	CrashLifetimeSecs interface{} `yaml:"crash_lifetime_secs,omitempty"`

	/*MemoryOvercommitFactor - Descr:  Default: 1
*/
	MemoryOvercommitFactor interface{} `yaml:"memory_overcommit_factor,omitempty"`

	/*StagingCpuLimitShares - Descr: CPU limit in shares for staging tasks cgroup Default: 512
*/
	StagingCpuLimitShares interface{} `yaml:"staging_cpu_limit_shares,omitempty"`

	/*InstanceDiskInodeLimit - Descr: Limit on inodes for an instance container Default: 200000
*/
	InstanceDiskInodeLimit interface{} `yaml:"instance_disk_inode_limit,omitempty"`

	/*ServerKey - Descr: PEM-encoded server key Default: <nil>
*/
	ServerKey interface{} `yaml:"server_key,omitempty"`

	/*CaCert - Descr: PEM-encoded CA certificate Default: <nil>
*/
	CaCert interface{} `yaml:"ca_cert,omitempty"`

	/*MaxStagingDuration - Descr:  Default: 900
*/
	MaxStagingDuration interface{} `yaml:"max_staging_duration,omitempty"`

}