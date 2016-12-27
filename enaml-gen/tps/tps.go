package tps 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Tps struct {

	/*TrafficControllerUrl - Descr: URL of Traffic controller Default: <nil>
*/
	TrafficControllerUrl interface{} `yaml:"traffic_controller_url,omitempty"`

	/*Cc - Descr: Basic auth username for CC internal API Default: internal_user
*/
	Cc *Cc `yaml:"cc,omitempty"`

	/*Watcher - Descr: address at which to serve debug info Default: 0.0.0.0:17015
*/
	Watcher *Watcher `yaml:"watcher,omitempty"`

	/*DropsondePort - Descr: local metron agent's port Default: 3457
*/
	DropsondePort interface{} `yaml:"dropsonde_port,omitempty"`

	/*ConsulAgentPort - Descr: local consul agent's port Default: 8500
*/
	ConsulAgentPort interface{} `yaml:"consul_agent_port,omitempty"`

	/*Bbs - Descr: maximum number of idle http connections Default: <nil>
*/
	Bbs *Bbs `yaml:"bbs,omitempty"`

	/*Listener - Descr: address at which to serve debug info Default: 0.0.0.0:17014
*/
	Listener *Listener `yaml:"listener,omitempty"`

	/*MaxInFlightRequests - Descr: Maximum number of requests to handle at once. Default: 200
*/
	MaxInFlightRequests interface{} `yaml:"max_in_flight_requests,omitempty"`

	/*LogLevel - Descr: Log level Default: info
*/
	LogLevel interface{} `yaml:"log_level,omitempty"`

}