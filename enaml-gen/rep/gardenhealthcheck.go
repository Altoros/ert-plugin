package rep 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type GardenHealthcheck struct {

	/*Interval - Descr: Frequency for healtchecking garden Default: 10m
*/
	Interval interface{} `yaml:"interval,omitempty"`

	/*Process - Descr: User to use while performing a container healthcheck Default: vcap
*/
	Process *Process `yaml:"process,omitempty"`

	/*CommandRetryPause - Descr: Time to wait between retrying garden commands Default: 5s
*/
	CommandRetryPause interface{} `yaml:"command_retry_pause,omitempty"`

	/*Timeout - Descr: Maximum allowed time for garden healthcheck Default: 10m
*/
	Timeout interface{} `yaml:"timeout,omitempty"`

}