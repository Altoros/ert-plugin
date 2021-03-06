package doppler 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type DopplerJob struct {

	/*Doppler - Descr: I/O Timeout on sinks Default: 0
*/
	Doppler *Doppler `yaml:"doppler,omitempty"`

	/*DopplerEndpoint - Descr: Shared secret used to verify cryptographically signed dropsonde messages Default: <nil>
*/
	DopplerEndpoint *DopplerEndpoint `yaml:"doppler_endpoint,omitempty"`

	/*Loggregator - Descr: Enable ssl for all communication with etcd Default: false
*/
	Loggregator *Loggregator `yaml:"loggregator,omitempty"`

	/*MetronEndpoint - Descr: The port used to emit dropsonde messages to the Metron agent Default: 3457
*/
	MetronEndpoint *MetronEndpoint `yaml:"metron_endpoint,omitempty"`

}