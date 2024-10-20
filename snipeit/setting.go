package snipeit

// Backup represents a configuration file backup.
type Backup struct {
	Filename      string `json:"filename"`
	Filesize      int    `json:"filesize"`
	ModifiedValue int64  `json:"modified_value"`
	BackupURL     string `json:"backup_url"`
}
