package goactor

import (
	"github.com/Colin1989/goactor/config"
	"github.com/Colin1989/goactor/logger"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/viper"
	"testing"
)

func TestAllOptions(t *testing.T) {
	// 启动worker
	vipNew := viper.New()
	conf := config.NewConfig(vipNew)

	opts := make([]Option, 0)
	// debug option
	opts = append(opts, WithDebug())
	// server mode, default Cluster
	opts = append(opts, WithSeverMode(Standalone))
	// node config
	vipNew.SetDefault("appConfig.NodeId", "1")
	vipNew.SetDefault("AppConfig.NodeType", "test")
	opts = append(opts, WithNodeConfig(conf))
	// logger config
	vipNew.SetDefault("logger.level", "error")
	vipNew.SetDefault("logger.LogPath", "./test_log")
	vipNew.SetDefault("logger.Maxsize", 1)
	vipNew.SetDefault("logger.maxAge", 1)
	vipNew.SetDefault("logger.MaxBackups", 100)
	opts = append(opts, WithLogger(conf))
	app := NewApp(opts...)

	assert.Equal(t, app.AppConfig, config.AppConfig{
		NodeId:   "1",
		NodeType: "test",
	})

	logger.Info("hello info")
	logger.Debug("hello debug")
	logger.Error("hello error")
}

func TestNewApp(t *testing.T) {
	tests := []struct {
		name string
		opts []Option
	}{
		// TODO: Add test cases.
		{name: "test-all", opts: []Option{
			WithDebug(),
			WithSeverMode(Standalone),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewApp(tt.opts...)
			if got == nil {
				return
			}
		})
	}
}
