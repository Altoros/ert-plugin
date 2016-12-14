package cloud_controller_ng 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Login struct {

	/*Enabled - Descr: whether use login as the authorization endpoint or not Default: true
*/
	Enabled interface{} `yaml:"enabled,omitempty"`

	/*Protocol - Descr: http or https Default: https
*/
	Protocol interface{} `yaml:"protocol,omitempty"`

	/*Url - Descr: URL of the login server Default: <nil>
*/
	Url interface{} `yaml:"url,omitempty"`

}