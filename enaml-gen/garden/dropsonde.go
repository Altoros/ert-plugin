package garden 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Dropsonde struct {

	/*Destination - Descr: A URL that points at the Metron agent to which metrics are forwarded. By default, it matches with the default of Metron. Default: <nil>
*/
	Destination interface{} `yaml:"destination,omitempty"`

	/*Origin - Descr: A string identifier that will be used when reporting metrics to Dropsonde. Default: <nil>
*/
	Origin interface{} `yaml:"origin,omitempty"`

}