package ssh_proxy 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Diego struct {

	/*Ssl - Descr: when connecting over https, ignore bad ssl certificates Default: false
*/
	Ssl *Ssl `yaml:"ssl,omitempty"`

	/*SshProxy - Descr: Comma separated list of allowed MAC algorithms Default: <nil>
*/
	SshProxy *SshProxy `yaml:"ssh_proxy,omitempty"`

}