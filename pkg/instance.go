package pkg

import (
	"sync"
)

var (
	tkwsSingleInstance *Manager
	once               sync.Once
)

// GetSingleInstance gets the singleton instance of WebSocket manager
func GetSingleInstance() *Manager {
	once.Do(func() {
		tkwsSingleInstance = NewManager()
		go tkwsSingleInstance.Start()
	})
	return tkwsSingleInstance
}

// NewInstance creates a new WebSocket manager instance (multi-instance mode)
func NewInstance() *Manager {
	tkws := NewManager()
	go tkws.Start()
	return tkws
}
