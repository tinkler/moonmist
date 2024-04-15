package internal_moonmist

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type GenConf struct {
	Dir   string         `yaml:"dir"`
	Codes []*GenCodeConf `yaml:"codes"`
}

type GenCodeConf struct {
	Type string `yaml:"type"`
	Out  string `yaml:"out"`
}

func GetGenConf(path string, root string) *GenConf {
	conf := &GenConf{}
	confFilePath := filepath.Join(root, path)
	fileData, err := os.ReadFile(confFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return defaultGenConf(root)
		}
		panic(err)
	}
	err = yaml.Unmarshal(fileData, conf)
	if err != nil {
		panic(err)
	}
	conf.Dir = filepath.Join(root, conf.Dir)
	for _, c := range conf.Codes {
		c.Out = filepath.Join(root, c.Out)
	}
	return conf
}

func defaultGenConf(root string) *GenConf {
	return &GenConf{
		Dir: filepath.Join(root, "./pkg/model"),
	}
}
