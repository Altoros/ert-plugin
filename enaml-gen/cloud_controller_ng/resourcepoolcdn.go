package cloud_controller_ng 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type ResourcePoolCdn struct {

	/*KeyPairId - Descr: Key pair name for signed download URIs Default: 
*/
	KeyPairId interface{} `yaml:"key_pair_id,omitempty"`

	/*PrivateKey - Descr: Private key for signing download URIs Default: 
*/
	PrivateKey interface{} `yaml:"private_key,omitempty"`

	/*Uri - Descr: URI for a CDN to used for resource pool downloads Default: 
*/
	Uri interface{} `yaml:"uri,omitempty"`

}