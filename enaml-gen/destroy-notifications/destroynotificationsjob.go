package destroy_notifications 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type DestroyNotificationsJob struct {

	/*Notifications - Descr: Username of the CF admin user Default: <nil>
*/
	Notifications *Notifications `yaml:"notifications,omitempty"`

	/*Ssl - Descr: Whether to verify SSL certs when making HTTP and SMTP requests Default: <nil>
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

	/*Domain - Descr: Cloud Foundry System Domain Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

}