package cloud_controller_worker 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type Droplets struct {

	/*Cdn - Descr: URI for a CDN to used for droplet downloads Default: 
*/
	Cdn *DropletsCdn `yaml:"cdn,omitempty"`

	/*FogConnection - Descr: Fog connection hash Default: <nil>
*/
	FogConnection interface{} `yaml:"fog_connection,omitempty"`

	/*WebdavConfig - Descr: The location of the webdav server eg: https://blobstore.internal Default: https://blobstore.service.cf.internal:4443
*/
	WebdavConfig *DropletsWebdavConfig `yaml:"webdav_config,omitempty"`

	/*FogAwsStorageOptions - Descr: Storage options passed to fog for aws blobstores. Valid keys: ['encryption']. Default: <nil>
*/
	FogAwsStorageOptions interface{} `yaml:"fog_aws_storage_options,omitempty"`

	/*BlobstoreType - Descr: The type of blobstore backing to use. Valid values: ['fog', 'webdav'] Default: fog
*/
	BlobstoreType interface{} `yaml:"blobstore_type,omitempty"`

	/*DropletDirectoryKey - Descr: Directory (bucket) used store droplets.  It does not have be pre-created. Default: cc-droplets
*/
	DropletDirectoryKey interface{} `yaml:"droplet_directory_key,omitempty"`

}