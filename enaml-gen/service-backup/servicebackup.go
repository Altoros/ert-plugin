package service_backup 
/*
* File Generated by enaml generator
* !!! Please do not edit this file !!!
*/
type ServiceBackup struct {

	/*ExitIfInProgress - Descr: Optional field to reject subsequent backup requests if a backup is already in progress. Defaults to false. Default: false
*/
	ExitIfInProgress interface{} `yaml:"exit_if_in_progress,omitempty"`

	/*MissingPropertiesMessage - Descr: Custom message to show when required backup parameters are not present. Default: Provide these missing fields in your manifest.
*/
	MissingPropertiesMessage interface{} `yaml:"missing_properties_message,omitempty"`

	/*ServiceIdentifierExecutable - Descr: Local executable to identify the service. If provided the identifier is included in log messages. Tokens are split on spaces; first is command to execute and remaining are passed as args to command. Default: <nil>
*/
	ServiceIdentifierExecutable interface{} `yaml:"service_identifier_executable,omitempty"`

	/*SourceExecutable - Descr: Local executable to create backups. Tokens are split on spaces; first is command to execute and remaining are passed as args to command. Default: <nil>
*/
	SourceExecutable interface{} `yaml:"source_executable,omitempty"`

	/*CronSchedule - Descr: Schedule on which to perform backups. Supports "* * * * *" syntax. Default: <nil>
*/
	CronSchedule interface{} `yaml:"cron_schedule,omitempty"`

	/*CleanupExecutable - Descr: Local executable to cleanup backups. Tokens are split on spaces; first is command to execute and remaining are passed as args to command. Default: 
*/
	CleanupExecutable interface{} `yaml:"cleanup_executable,omitempty"`

	/*SourceFolder - Descr: Local path from which backups are uploaded. All files in here are uploaded. Default: <nil>
*/
	SourceFolder interface{} `yaml:"source_folder,omitempty"`

	/*Destination - Descr: SSH port on remote server Default: 22
*/
	Destination *Destination `yaml:"destination,omitempty"`

}