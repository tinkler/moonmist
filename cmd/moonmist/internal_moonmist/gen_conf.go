package internal_moonmist

import (
	"errors"
	"os"

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

func GetGenConf(path string) *GenConf {
	conf := &GenConf{}
	fileData, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return defaultGenConf()
		}
		panic(err)
	}
	err = yaml.Unmarshal(fileData, conf)
	if err != nil {
		panic(err)
	}
	return conf
}

func defaultGenConf() *GenConf {
	return &GenConf{
		Dir: "./pkg/model",
	}
}
