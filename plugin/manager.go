package plugin

import (
	"fmt"
	goplugin "plugin"
)

// Manager is in charge of instanciating plugins and storing those in memory
type Manager struct {
	middlewares []*Middleware
}

// NewManager builds a new manager
func NewManager() *Manager {
	return &Manager{
		middlewares: []*Middleware{},
	}
}

// Load loads a plugin
func (m *Manager) Load(plugin *Plugin) error {
	p, err := goplugin.Open(plugin.Path)
	if err != nil {
		return fmt.Errorf("Error opening plugin: %s", err)
	}
	loader, err := p.Lookup("Load")
	if err != nil {
		return fmt.Errorf("Error in plugin Lookup: %s", err)
	}
	instance := loader.(func() interface{})()
	middleware, ok := instance.(Middleware)
	if ok {
		// if Middleware, add to middleware list
		m.middlewares = append(m.middlewares, &middleware)
	} else {
		return fmt.Errorf("Plugin from %+v does not implement any plugin interface", plugin)
	}
	return nil
}

// GetMiddlewares return a list of all Middleware plugins
func (m *Manager) GetMiddlewares() []*Middleware {
	return m.middlewares
}
