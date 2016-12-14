package test_notifications_ui 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type NotificationsUi struct {

	/*Space - Descr: Notifications space Default: <nil>
*/
	Space interface{} `yaml:"space,omitempty"`

	/*Cf - Descr: Password of the CF admin user Default: <nil>
*/
	Cf *Cf `yaml:"cf,omitempty"`

	/*AppDomain - Descr: Domain used to host application Default: <nil>
*/
	AppDomain interface{} `yaml:"app_domain,omitempty"`

	/*Uaa - Descr: UAA Admin client secret Default: <nil>
*/
	Uaa *Uaa `yaml:"uaa,omitempty"`

	/*Organization - Descr: Notifications organization Default: <nil>
*/
	Organization interface{} `yaml:"organization,omitempty"`

}