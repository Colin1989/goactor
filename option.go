package goactor

import (
	"fmt"
	"github.com/Colin1989/goactor/config"
	"github.com/Colin1989/goactor/logger"
	"reflect"
)

// Option is a function on the options for a connection.
type Option func(*App) error

func WithNodeConfig(conf *config.Config) Option {
	return func(app *App) error {
		name := reflect.TypeOf(app.AppConfig).Name()
		if err := conf.UnmarshalKey(name, &app.AppConfig); err != nil {
			return err
		}
		fmt.Println(conf)
		return nil
	}
}

// WithLogger
//
//	@Description: 需要在 WithNodeConfig 之后调用
//	@param conf
//	@return Option
func WithLogger(conf *config.Config) Option {
	return func(app *App) error {
		logger.SetNodeLogger(app.AppConfig, conf)
		return nil
	}
}

func WithDebug() Option {
	return func(a *App) error {
		a.debug = true
		return nil
	}
}

func WithSeverMode(mode ServerMode) Option {
	return func(a *App) error {
		a.serverMode = mode
		return nil
	}
}

func WithClusterMode() Option {
	return func(app *App) error {
		return nil
	}
}
