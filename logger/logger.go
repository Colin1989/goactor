package logger

import (
	"github.com/Colin1989/goactor/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"strings"
)

var (
	DefaultLogger *zap.SugaredLogger
)

func init() {
	conf2 := newDefaultLogger()
	DefaultLogger = newLogger(config.NewDefaultAppConfig(), conf2)
}

func SetNodeLogger(app config.AppConfig, conf *config.Config) *zap.SugaredLogger {
	conf2 := NewLoggerConfig(conf)
	DefaultLogger = newLogger(app, conf2)
	return DefaultLogger
}

func newLogger(app config.AppConfig, conf Logger) *zap.SugaredLogger {
	core := zapcore.NewTee(normalCore(app.String(), conf), errorCore(app.String(), conf))
	opt := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	}
	field := zap.String("node", app.String())
	return zap.New(core, opt...).With(field).Sugar()
}

func normalCore(fileName string, opt Logger) zapcore.Core {
	//获取编码器
	encoderConfig := zap.NewProductionEncoderConfig()     //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder //指定时间格式
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder      	//显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	optLevel := GetLevel(opt.Level)
	normalPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= optLevel
	})

	writeSyncers := make([]zapcore.WriteSyncer, 0)
	normalFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path.Join(opt.LogPath, fileName),
		MaxSize:    opt.MaxSize,
		MaxAge:     opt.MaxAge,
		MaxBackups: opt.MaxBackups,
		Compress:   true,
	})
	writeSyncers = append(writeSyncers, normalFileWriteSyncer)
	writeSyncers = append(writeSyncers, zapcore.Lock(os.Stdout))

	normalFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncers...), normalPriority)
	return normalFileCore
}

// errorCore
//
//	@Description: 生成 error zap
//	@param opt
//	@return zapcore.Core
func errorCore(fileName string, opt Logger) zapcore.Core {
	//获取编码器
	encoderConfig := zap.NewProductionEncoderConfig()     //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder //指定时间格式
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder      	//显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	errPriority := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zap.ErrorLevel
	})

	errFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   path.Join(opt.LogPath, fileName+"_error"),
		MaxSize:    opt.MaxSize,
		MaxAge:     opt.MaxAge,
		MaxBackups: opt.MaxBackups,
		Compress:   true,
	})
	errFileCore := zapcore.NewCore(encoder, zapcore.AddSync(errFileWriteSyncer), errPriority)
	return errFileCore
}

func GetLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
