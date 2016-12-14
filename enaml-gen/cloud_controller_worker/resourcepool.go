package cloud_controller_worker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type ResourcePool struct {

	/*Cdn - Descr: Private key for signing download URIs Default: 
*/
	Cdn *ResourcePoolCdn `yaml:"cdn,omitempty"`

	/*ResourceDirectoryKey - Descr: Directory (bucket) used store app resources.  It does not have be pre-created. Default: cc-resources
*/
	ResourceDirectoryKey interface{} `yaml:"resource_directory_key,omitempty"`

	/*FogConnection - Descr: Fog connection hash Default: <nil>
*/
	FogConnection interface{} `yaml:"fog_connection,omitempty"`

	/*MinimumSize - Descr: Minimum size of a resource to add to the pool Default: 65536
*/
	MinimumSize interface{} `yaml:"minimum_size,omitempty"`

	/*FogAwsStorageOptions - Descr: Storage options passed to fog for aws blobstores. Valid keys: ['encryption']. Default: <nil>
*/
	FogAwsStorageOptions interface{} `yaml:"fog_aws_storage_options,omitempty"`

	/*MaximumSize - Descr: Maximum size of a resource to add to the pool Default: 536870912
*/
	MaximumSize interface{} `yaml:"maximum_size,omitempty"`

	/*BlobstoreType - Descr: The type of blobstore backing to use. Valid values: ['fog', 'webdav'] Default: fog
*/
	BlobstoreType interface{} `yaml:"blobstore_type,omitempty"`

	/*WebdavConfig - Descr: The basic auth user that CC uses to connect to the admin endpoint on webdav Default: 
*/
	WebdavConfig *ResourcePoolWebdavConfig `yaml:"webdav_config,omitempty"`

}