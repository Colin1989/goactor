package goactor

import (
	"github.com/Colin1989/goactor/config"
)

// ServerMode represents a server mode
type ServerMode byte

const (
	_ ServerMode = iota
	// Cluster represents a server running with connection to other servers
	Cluster
	// Standalone represents a server running without connection to other servers
	Standalone
)

type App struct {
	config.AppConfig
	debug      bool
	dieChan    chan struct{}
	serverMode ServerMode
}

func NewApp(opts ...Option) *App {
	app := &App{
		debug:      false,
		dieChan:    make(chan struct{}),
		serverMode: Cluster,
	}

	app.AppConfig = config.NewDefaultAppConfig()
	//logger.SetNodeLogger("node")

	for _, opt := range opts {
		if err := opt(app); err != nil {
			panic(err)
		}
	}

	return app
}
