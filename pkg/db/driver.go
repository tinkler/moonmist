package db

import (
	"sync"

	"gorm.io/gorm"
)

var (
	supportDrivers   map[string]func(dsn string) (*gorm.DB, error)
	supportDriversMu sync.Mutex
)

func RegisterDriver(driver string, open func(dsn string) (*gorm.DB, error)) {
	supportDriversMu.Lock()
	defer supportDriversMu.Unlock()
	_, duplicated := supportDrivers[driver]
	if duplicated {
		return
	}
	supportDrivers[driver] = open
}
