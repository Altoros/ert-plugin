package blobstore 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Blobstore struct {

	/*Tls - Descr: The PEM-encoded private key for signing TLS/SSL traffic Default: <nil>
*/
	Tls *Tls `yaml:"tls,omitempty"`

	/*SecureLink - Descr: The secret used for signing URLs Default: <nil>
*/
	SecureLink *SecureLink `yaml:"secure_link,omitempty"`

	/*AdminUsers - Descr: List of Username and Password pairs that have admin access to the blobstore. Cloud Controller must use one of these to access the blobstore via HTTP Basic Auth.
Example:
  users:
  - username: user1
    password: password1
  - username: user2
    password: password2
 Default: <nil>
*/
	AdminUsers interface{} `yaml:"admin_users,omitempty"`

	/*MaxUploadSize - Descr: Max allowed file size for upload Default: 5000m
*/
	MaxUploadSize interface{} `yaml:"max_upload_size,omitempty"`

	/*NginxWorkersPerCore - Descr: Number of NGINX worker processes per CPU core Default: 2
*/
	NginxWorkersPerCore interface{} `yaml:"nginx_workers_per_core,omitempty"`

	/*InternalAccessRules - Descr: List of allow / deny rules for the blobstore internal server. Defaults to RFC 1918 Private Networks. Will be followed by 'deny all'. See http://nginx.org/en/docs/http/ngx_http_access_module.html for valid rules Default: [allow 10.0.0.0/8; allow 172.16.0.0/12; allow 192.168.0.0/16;]
*/
	InternalAccessRules interface{} `yaml:"internal_access_rules,omitempty"`

	/*Port - Descr: TCP port on which the blobstore server (nginx) listens Default: 8080
*/
	Port interface{} `yaml:"port,omitempty"`

}