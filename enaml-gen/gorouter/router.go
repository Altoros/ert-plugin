package gorouter 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Router struct {

	/*DrainWait - Descr: Delay in seconds after drain begins before server stops listening.
During this time the server will respond with 503 Service Unavailable to
requests having header User-Agent: HTTP-Monitor/1.1. This accommodates
requests in transit sent while health check responded ok.
 Default: 0
*/
	DrainWait interface{} `yaml:"drain_wait,omitempty"`

	/*CipherSuites - Descr: An ordered list of supported SSL cipher suites containing golang tls constants separated by colons The cipher suite will be chosen according to this order during SSL handshake Default: TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA:TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA:TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA:TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA:TLS_RSA_WITH_AES_128_CBC_SHA:TLS_RSA_WITH_AES_256_CBC_SHA
*/
	CipherSuites interface{} `yaml:"cipher_suites,omitempty"`

	/*RouteServicesSecretDecryptOnly - Descr: To rotate keys, add your new key here and deploy. Then swap this key with the value of route_services_secret and deploy again. Default: 
*/
	RouteServicesSecretDecryptOnly interface{} `yaml:"route_services_secret_decrypt_only,omitempty"`

	/*SslKey - Descr: The private ssl key for ssl termination Default: 
*/
	SslKey interface{} `yaml:"ssl_key,omitempty"`

	/*RouteServicesSecret - Descr: Support for route services is disabled when no value is configured. A robust passphrase is recommended. Default: 
*/
	RouteServicesSecret interface{} `yaml:"route_services_secret,omitempty"`

	/*Offset - Descr:  Default: 0
*/
	Offset interface{} `yaml:"offset,omitempty"`

	/*DnsHealthCheckHost - Descr: Host to ping for confirmation of DNS resolution, only used when Routing API is enabled Default: consul.service.cf.internal
*/
	DnsHealthCheckHost interface{} `yaml:"dns_health_check_host,omitempty"`

	/*SslSkipValidation - Descr: Skip SSL client cert validation Default: false
*/
	SslSkipValidation interface{} `yaml:"ssl_skip_validation,omitempty"`

	/*Port - Descr: Listening Port for Router. Default: 80
*/
	Port interface{} `yaml:"port,omitempty"`

	/*SslCert - Descr: The public ssl cert for ssl termination Default: 
*/
	SslCert interface{} `yaml:"ssl_cert,omitempty"`

	/*RequestedRouteRegistrationIntervalInSeconds - Descr: On startup, the router will delay listening for requests by this duration to increase likelihood that it has a complete routing table before serving requests. The router also broadcasts the same duration as a recommended interval to registering clients via NATS. Default: 20
*/
	RequestedRouteRegistrationIntervalInSeconds interface{} `yaml:"requested_route_registration_interval_in_seconds,omitempty"`

	/*DebugAddress - Descr: Address at which to serve debug info Default: 0.0.0.0:17002
*/
	DebugAddress interface{} `yaml:"debug_address,omitempty"`

	/*TraceKey - Descr: If the X-Vcap-Trace request header is set and has this value, trace headers are added to the response. Default: 22
*/
	TraceKey interface{} `yaml:"trace_key,omitempty"`

	/*EnableProxy - Descr: Enables support for the popular PROXY protocol, allowing downstream load balancers that do not support HTTP to pass along client information. Default: false
*/
	EnableProxy interface{} `yaml:"enable_proxy,omitempty"`

	/*Logrotate - Descr: The number of files that logrotate will keep around on the VM Default: 7
*/
	Logrotate *Logrotate `yaml:"logrotate,omitempty"`

	/*LoggingLevel - Descr: Log level for router Default: info
*/
	LoggingLevel interface{} `yaml:"logging_level,omitempty"`

	/*EnableSsl - Descr: Enable ssl termination on the router Default: false
*/
	EnableSsl interface{} `yaml:"enable_ssl,omitempty"`

	/*RouteServicesTimeout - Descr: Expiry time of a route service signature in seconds Default: 60
*/
	RouteServicesTimeout interface{} `yaml:"route_services_timeout,omitempty"`

	/*ExtraHeadersToLog - Descr: A list of headers that log events will be annotated with Default: []
*/
	ExtraHeadersToLog interface{} `yaml:"extra_headers_to_log,omitempty"`

	/*NumberOfCpus - Descr: Number of CPUs to utilize, the default (-1) will equal the number of available CPUs Default: -1
*/
	NumberOfCpus interface{} `yaml:"number_of_cpus,omitempty"`

	/*EnableAccessLogStreaming - Descr: Enables streaming of access log to syslog. Warning: this comes with a performance cost; due to higher I/O, max request rate is reduced. Default: false
*/
	EnableAccessLogStreaming interface{} `yaml:"enable_access_log_streaming,omitempty"`

	/*SecureCookies - Descr: Set secure flag on http cookies Default: false
*/
	SecureCookies interface{} `yaml:"secure_cookies,omitempty"`

	/*Status - Descr: Username for HTTP basic auth to the varz/status endpoint. Default: <nil>
*/
	Status *Status `yaml:"status,omitempty"`

	/*RouteServicesRecommendHttps - Descr: Route Services are told where to send requests after processing using the X-CF-Forwarded-Url header. When this property is true, the scheme for this URL is https. When false, the scheme is http. As requests from Route Services to applications on CF transit load balancers and gorouter, disable this property for deployments that have TLS termination disabled. Default: true
*/
	RouteServicesRecommendHttps interface{} `yaml:"route_services_recommend_https,omitempty"`

}