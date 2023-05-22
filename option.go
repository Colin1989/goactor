package goactor

// Option is a function on the options for a connection.
type Option func(*App) error

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
