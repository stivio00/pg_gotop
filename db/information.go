package db

type DbCurrentConnection struct {
	User     string
	Database string
	Version  string
	Size     string
}

type DbInformation struct {
	Databases int
	Tables    int
	Schemas   int
	Size      uint64
}

type DbSettings struct {
	SharedBufferSize   string
	TempBuffers        string
	MaintenanceWorkMem string
	WorkingMem         string
	EffectiveCacheSize string
	ListenAddresses    string
	MaxConnections     string
	AutovacuumEnabled  string
	Fillfactor         string
}

type FileLocations struct {
	Data_directory    string
	Config_file       string
	Hba_file          string
	Ident_file        string
	External_pid_file string
}

func (d *DbConnection) GetDbInfo() DbInformation {
	return DbInformation{
		Databases: 4,
		Tables:    1,
		Schemas:   2,
		Size:      10101,
	}
}

func (d *DbConnection) GetDbSettings() DbSettings {
	return DbSettings{
		SharedBufferSize:   d.GetSetting("shared_buffers"),
		TempBuffers:        d.GetSetting("temp_buffers"),
		MaintenanceWorkMem: d.GetSetting("maintenance_work_mem"),
		WorkingMem:         d.GetSetting("work_mem"),
		EffectiveCacheSize: d.GetSetting("effective_cache_size"),
		ListenAddresses:    d.GetSetting("listen_addresses"),
		MaxConnections:     d.GetSetting("max_connections"),
		AutovacuumEnabled:  d.GetSetting("autovacuum_enabled"),
		Fillfactor:         d.GetSetting("fillfactor"),
	}
}

func (d *DbConnection) GetFileLocations() FileLocations {
	return FileLocations{
		Data_directory:    d.GetSetting("data_directory"),
		Config_file:       d.GetSetting("config_file"),
		Hba_file:          d.GetSetting("hba_file"),
		Ident_file:        d.GetSetting("ident_file"),
		External_pid_file: d.GetSetting("external_pid_file"),
	}
}

func (d *DbConnection) GetDbCurrentConnection() DbCurrentConnection {
	current := DbCurrentConnection{}
	v := "(SELECT pg_size_pretty(sum(pg_tablespace_size(oid))) AS \"total cluster size\" FROM pg_tablespace) \"size\""
	row := d.db.QueryRow("SELECT current_user, current_catalog, version(), " + v + " ; ")
	row.Scan(&current.User, &current.Database, &current.Version, &current.Size)
	return current
}

func (d *DbConnection) GetSetting(settingName string) string {
	var value string
	row := d.db.QueryRow("SELECT current_setting ('" + settingName + "');")
	row.Scan(&value)
	return value
}
