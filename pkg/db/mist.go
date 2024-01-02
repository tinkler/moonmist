package db

import (
	"context"
	"sync"

	"github.com/tinkler/moonmist/pkg/mist"
	"gorm.io/gorm"
)

type contextKey string

const (
	mistKey          contextKey = "mist"
	defaultGroupName            = "default"
	MiddlewareName              = "db"
)

func init() {
	mist.Register[*gorm.DB](MiddlewareName, &Middleware{})
}

type MiddlewareConfig struct {
	Configs []*Config
}

func (c MiddlewareConfig) GetName() string {
	return MiddlewareName
}

type Middleware struct {
	dms     map[string]*DatabaseManager
	configs []*Config
}

func (s *Middleware) Config() interface{} {
	return &s.configs
}

func (m *Middleware) Run(ctx context.Context) error {
	m.MustSupport(m.getAllDrivers())
	m.dms = make(map[string]*DatabaseManager)
	var defaultDatabaseManager *DatabaseManager
	for i, c := range m.configs {
		_, duplicated := m.dms[c.Name]
		if duplicated {
			panic("duplicated database config")
		}
		dm := new(DatabaseManager)
		dm.config = c
		m.dms[c.Name] = dm
		if i == 0 {
			defaultDatabaseManager = dm
		}
	}
	m.dms["default"] = defaultDatabaseManager
	return nil
}

func (m *Middleware) getAllDrivers() []string {
	driverMap := make(map[string]bool)
	for _, config := range m.configs {
		for _, driverConfig := range config.Drivers {
			driverMap[driverConfig.Driver] = true
		}
	}
	var drivers []string
	for driver := range driverMap {
		drivers = append(drivers, driver)
	}
	return drivers
}

func (m *Middleware) MustSupport(drivers []string) {
	for _, driver := range drivers {
		_, supported := supportDrivers[driver]
		if !supported {
			panic("driver " + driver + " is not supported, forget to Register?")
		}
	}
}

type DatabaseManager struct {
	init   sync.Once
	config *Config
	dbm    map[string]*gorm.DB
}

func (dm *DatabaseManager) Of(ctx context.Context, args ...interface{}) *gorm.DB {
	dm.init.Do(func() {
		var defaultDB *gorm.DB
		for i, driverConfig := range dm.config.Drivers {
			newDB, err := supportDrivers[driverConfig.Driver](driverConfig.Dsn)
			if err != nil {
				panic(err)
			}
			if i == 0 {
				defaultDB = newDB
			}
			dm.dbm[driverConfig.Name] = newDB
		}
		dm.dbm[dm.config.DefaultDriver] = defaultDB
	})
	return dm.dbm[dm.config.DefaultDriver]
}

func (m *Middleware) Of(ctx context.Context, args ...interface{}) *gorm.DB {
	groupName := defaultGroupName
	if len(args) > 0 {
		if n, ok := args[0].(string); ok {
			groupName = n
		}
	}
	if db := m.dms[groupName]; db != nil {
		return db.Of(ctx, args[1:]...)
	}
	panic("can not found db")
}

type Config struct {
	Name          string    `yaml:"name"`
	DefaultDriver string    `yaml:"default_driver_name"`
	Drivers       []*Driver `yaml:"drivers"`
}

func (c Config) GetName() string {
	return c.Name
}

type Driver struct {
	Name   string `yaml:"name" validate:"required"`
	Driver string `yaml:"driver" validate:"required"`
	Dsn    string `yaml:"dsn" validate:"required"`
}
