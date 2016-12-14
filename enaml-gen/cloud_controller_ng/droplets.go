package cloud_controller_ng 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Droplets struct {

	/*DropletDirectoryKey - Descr: Directory (bucket) used store droplets.  It does not have be pre-created. Default: cc-droplets
*/
	DropletDirectoryKey interface{} `yaml:"droplet_directory_key,omitempty"`

	/*WebdavConfig - Descr: The basic auth user that CC uses to connect to the admin endpoint on webdav Default: 
*/
	WebdavConfig *DropletsWebdavConfig `yaml:"webdav_config,omitempty"`

	/*Cdn - Descr: Private key for signing download URIs Default: 
*/
	Cdn *DropletsCdn `yaml:"cdn,omitempty"`

	/*MaxStagedDropletsStored - Descr: Number of recent, staged droplets stored per app (not including current droplet) Default: 5
*/
	MaxStagedDropletsStored interface{} `yaml:"max_staged_droplets_stored,omitempty"`

	/*BlobstoreType - Descr: The type of blobstore backing to use. Valid values: ['fog', 'webdav'] Default: fog
*/
	BlobstoreType interface{} `yaml:"blobstore_type,omitempty"`

	/*FogConnection - Descr: Fog connection hash Default: <nil>
*/
	FogConnection interface{} `yaml:"fog_connection,omitempty"`

	/*FogAwsStorageOptions - Descr: Storage options passed to fog for aws blobstores. Valid keys: ['encryption']. Default: <nil>
*/
	FogAwsStorageOptions interface{} `yaml:"fog_aws_storage_options,omitempty"`

}