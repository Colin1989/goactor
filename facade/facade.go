package facade

// Module is the interface that represent a module.
type Module interface {
	Init() error
	AfterInit()
	BeforeShutdown()
	Shutdown() error
}
