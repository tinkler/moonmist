package internal_moonmist

import (
	"os"

	"github.com/tinkler/moonmist/pkg/runtime"
)

func MkdirAll(conf *GenConf) {
	if conf == nil {
		panic("nil conf")
	}
	for _, code := range conf.Codes {
		if code.Out == "" {
			continue
		}
		runtime.Must(os.MkdirAll(code.Out, os.ModePerm))
	}
}
