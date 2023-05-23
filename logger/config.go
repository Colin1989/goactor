package logger

import (
	"fmt"
	"github.com/Colin1989/goactor/config"
	"reflect"
)

type Logger struct {
	Level      string `json:"level"`       // 输出日志等级
	LogPath    string `json:"log_path"`    // 日志保存路径
	MaxSize    int    `json:"max_size"`    // 文件切割大小(MB)
	MaxAge     int    `json:"max_age"`     // 最大保留天数(达到限制，则会被清理)
	MaxBackups int    `json:"max_backups"` // 最大备份数量
}

func newDefaultLogger() Logger {
	return Logger{
		Level:      "debug",
		LogPath:    "./log",
		MaxSize:    10,
		MaxAge:     30,
		MaxBackups: 100,
	}
}

func NewLoggerConfig(conf *config.Config) Logger {
	conf2 := newDefaultLogger()
	name := reflect.TypeOf(conf2).Name()
	if err := conf.UnmarshalKey(name, &conf2); err != nil {
		panic(err)
	}
	fmt.Println(conf2)
	return conf2
}
