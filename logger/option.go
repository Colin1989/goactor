package logger

type Option struct {
	Level           string `json:"level"`             // 输出日志等级
	EnableConsole   bool   `json:"enable_console"`    // 是否控制台输出
	EnableWriteFile bool   `json:"enable_write_file"` // 是否输出文件(必需配置FilePath)
	LogPath         string `json:"log_path"`          // 日志保存路径
	FileName        string `json:"file_name"`         // 日志文件名称
	MaxSize         int    `json:"max_size"`          // 文件切割大小(MB)
	MaxAge          int    `json:"max_age"`           // 最大保留天数(达到限制，则会被清理)
	MaxBackups      int    `json:"max_backups"`       // 最大备份数量
}
