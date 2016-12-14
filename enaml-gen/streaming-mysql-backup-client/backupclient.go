package streaming_mysql_backup_client 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type BackupClient struct {

	/*OutputFolder - Descr: Folder to place the prepared backups Default: /var/vcap/store/mysql-backups
*/
	OutputFolder interface{} `yaml:"output_folder,omitempty"`

	/*TmpFolder - Descr: Folder to download / prepare backups Default: /var/vcap/store/mysql-backups-tmp
*/
	TmpFolder interface{} `yaml:"tmp_folder,omitempty"`

}