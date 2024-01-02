package mist

import (
	"context"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"unsafe"

	"gopkg.in/yaml.v3"
)

var (
	mists             = make(map[string]Middleware)
	typeMap           = make(map[unsafe.Pointer]string)
	mu                sync.RWMutex
	defaultMistConfig = &MistConfig{}
)

type Middleware interface {
	Config() interface{}
}

type MistConfig struct {
	Middlewares []*MiddlewareConfig `yaml:"middlewares"`
}

type MiddlewareConfig struct {
	Name    string        `yaml:"name"`
	Configs []interface{} `yaml:"configs"`
}

type DriverCreater[C any] func(mistName string) error

func Register[T any](mistName string, driver Middleware) error {
	if driver == nil {
		panic("driver can not be nil")
	}
	mu.Lock()
	defer mu.Unlock()

	_, duplicated := mists[mistName]
	if duplicated {
		panic("duplicated register mist driver " + mistName)
	}
	mists[mistName] = driver
	var t T
	typeMap[unpackEFace(t).rtype] = mistName
	defaultMistConfig.Middlewares = append(defaultMistConfig.Middlewares, &MiddlewareConfig{
		Name: mistName,
	})
	return nil
}

func Get[T any](ctx context.Context) T {
	var t T
	expectRType := unpackEFace(t).rtype
	m, had := mists[typeMap[expectRType]]
	if !had {
		expectType := reflect.TypeOf(0)
		(*iface)(unsafe.Pointer(&expectType)).data = expectRType
		panic("driver type " + expectType.String() + " not found")
	}
	return m.(T)
}

func Run(ctx context.Context) error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	configData, err := os.ReadFile(filepath.Join(root, "mist.yaml"))
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(configData, defaultMistConfig)
	if err != nil {
		return err
	}
	for _, mc := range defaultMistConfig.Middlewares {
		if m, ok := mists[mc.Name]; ok {
			configData, err := json.Marshal(mc.Configs)
			if err != nil {
				return err
			}
			err = json.Unmarshal(configData, m.Config())
			if err != nil {
				return err
			}
		}
	}
	return nil
}
