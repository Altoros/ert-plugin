package doppler 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Loggregator struct {

	/*Etcd - Descr: Enable ssl for all communication with etcd Default: false
*/
	Etcd *LoggregatorEtcd `yaml:"etcd,omitempty"`

	/*Tls - Descr: CA root required for key/cert verification Default: 
*/
	Tls *LoggregatorTls `yaml:"tls,omitempty"`

}