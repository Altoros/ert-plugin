package uaa 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type UaaUser struct {

	/*Authorities - Descr: Contains a list of the default authorities/scopes assigned to a user Default: [openid scim.me cloud_controller.read cloud_controller.write cloud_controller_service_permissions.read password.write uaa.user approvals.me oauth.approvals notification_preferences.read notification_preferences.write profile roles user_attributes uaa.offline_token]
*/
	Authorities interface{} `yaml:"authorities,omitempty"`

}