package setting

type Config struct {
	Server   ServerSetting   `mapstructure:"server"`
	Database MySQLSetting    `mapstructure:"database"`
	Logger   LoggerSetting   `mapstructure:"logger"`
	Electron ElectronSetting `mapstructure:"electron"`
	JWT      JWTSetting      `mapstructure:"jwt"`
	Path     PathSetting     `mapstructure:"path"`
}
type ServerSetting struct {
	Port    string `mapstructure:"port"`
	Mode    string `mapstructure:"mode"`
	AppName string `mapstructure:"app_name"`
	Path    string `mapstructure:"path"`
}
type JWTSetting struct {
	Secret    string `mapstructure:"secret"`
	Durations int64  `mapstructure:"durations"`
}
type MySQLSetting struct {
	Path          string `mapstructure:"path"`
	PathMigration string `mapstructure:"path"`
	DbName        string `mapstructure:"db_name"`
}
type ElectronSetting struct {
	UserDataPath string
	Version      string
}
type LoggerSetting struct {
	Level       string `mapstructure:"level"`
	PathLog     string `mapstructure:"path_log"`
	FileLogName string `mapstructure:"file_log_name"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxBackups  int    `mapstructure:"max_backups"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
}

type PathSetting struct {
	AppDataDir  string `mapstructure:"app_data_dir"`
	AppDir      string `mapstructure:"app_dir"`
	PathStorage string `mapstructure:"path_storage"`
	PathAvatar  string `mapstructure:"path_avatar"`
	PathImage   string `mapstructure:"path_image"`
	PathFile    string `mapstructure:"path_file"`
	PathWeb     string `mapstructure:"path_web"`
}
