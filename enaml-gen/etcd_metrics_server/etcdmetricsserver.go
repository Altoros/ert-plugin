package etcd_metrics_server 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type EtcdMetricsServer struct {

	/*Etcd - Descr: DNS suffix of the etcd server to instrument.
Target etcd server must be colocated with this etcd_metrics_server.
This property is only used if 'etcd_metrics_server.etcd.require_ssl' is 'true'."
 Default: <nil>
*/
	Etcd *Etcd `yaml:"etcd,omitempty"`

	/*Status - Descr: basic auth password for metrics server (leave empty for generated) Default: 
*/
	Status *Status `yaml:"status,omitempty"`

	/*Nats - Descr: NATS server password Default: <nil>
*/
	Nats *Nats `yaml:"nats,omitempty"`

}