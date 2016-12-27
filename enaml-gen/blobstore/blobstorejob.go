package blobstore 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type BlobstoreJob struct {

	/*Blobstore - Descr: Number of NGINX worker processes per CPU core Default: 2
*/
	Blobstore *Blobstore `yaml:"blobstore,omitempty"`

	/*SystemDomain - Descr: The system domain.  The public server will listen on host 'blobstore.system-domain.tld' Default: <nil>
*/
	SystemDomain interface{} `yaml:"system_domain,omitempty"`

	/*Domain - Descr: DEPRECATED: The system domain.  The public server will listen on host 'blobstore.system-domain.tld' Default: <nil>
*/
	Domain interface{} `yaml:"domain,omitempty"`

}