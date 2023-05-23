package config

import (
	"github.com/spf13/viper"
	"strings"
	"time"
)

// Config is a wrapper around a viper config
type Config struct {
	config *viper.Viper
}

func NewConfig(cfgs ...*viper.Viper) *Config {
	var cfg *viper.Viper
	if len(cfgs) > 0 {
		cfg = cfgs[0]
	} else {
		cfg = viper.New()
	}

	// 首先点号('.')是不能出现在环境变量的，
	// 所以必须用viper.SetEnvKeyReplacer把点号('.')映射成下划线('_')，
	cfg.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	cfg.AutomaticEnv()
	c := &Config{config: cfg}
	c.fillDefaultValues()
	return c
}

func (c *Config) fillDefaultValues() {
	appConfig := NewDefaultAppConfig()

	defaultsMap := map[string]interface{}{
		"nodeId":   appConfig.NodeId,
		"nodeType": appConfig.NodeType,
	}

	for param := range defaultsMap {
		if c.config.Get(param) == nil {
			c.config.SetDefault(param, defaultsMap[param])
		}
	}
}

// GetDuration returns a duration from the inner config
func (c *Config) GetDuration(s string) time.Duration {
	return c.config.GetDuration(s)
}

// GetString returns a string from the inner config
func (c *Config) GetString(s string) string {
	return c.config.GetString(s)
}

// GetInt returns an int from the inner config
func (c *Config) GetInt(s string) int {
	return c.config.GetInt(s)
}

// GetBool returns an boolean from the inner config
func (c *Config) GetBool(s string) bool {
	return c.config.GetBool(s)
}

// GetStringSlice returns a string slice from the inner config
func (c *Config) GetStringSlice(s string) []string {
	return c.config.GetStringSlice(s)
}

// Get returns an interface from the inner config
func (c *Config) Get(s string) interface{} {
	return c.config.Get(s)
}

// GetStringMapString returns a string map string from the inner config
func (c *Config) GetStringMapString(s string) map[string]string {
	return c.config.GetStringMapString(s)
}

// UnmarshalKey unmarshals key into v
func (c *Config) UnmarshalKey(s string, v interface{}) error {
	return c.config.UnmarshalKey(s, v)
}

// Unmarshal unmarshals config into v
func (c *Config) Unmarshal(v interface{}) error {
	return c.config.Unmarshal(v)
}
