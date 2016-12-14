package cloud_controller_clock 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Api struct {

	/*RestartIfAboveMb - Descr: The cc will restart if memory remains above this threshold for 3 monit cycles Default: 3750
*/
	RestartIfAboveMb interface{} `yaml:"restart_if_above_mb,omitempty"`

	/*RestartIfConsistentlyAboveMb - Descr: The cc will restart if memory remains above this threshold for 15 monit cycles Default: 3500
*/
	RestartIfConsistentlyAboveMb interface{} `yaml:"restart_if_consistently_above_mb,omitempty"`

	/*AlertIfAboveMb - Descr: The cc will alert if memory remains above this threshold for 3 monit cycles Default: 3500
*/
	AlertIfAboveMb interface{} `yaml:"alert_if_above_mb,omitempty"`

}