package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tinkler/moonmist/cmd/moonmist/internal_moonmist"
	"github.com/tinkler/moonmist/internal/parser"
	"github.com/tinkler/moonmist/pkg/runtime"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

const (
	versionFlag = "version"
	fileFlag    = "file"
)

func parseEnvFlags() {
	pflag.Bool(versionFlag, false, "print version and exit")
	pflag.StringP(fileFlag, "f", "mist.yaml", "moonmist config file")
	pflag.Parse()

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	runtime.Must(viper.BindPFlags(pflag.CommandLine))

}

func main() {
	parseEnvFlags()
	if viper.GetBool(versionFlag) {
		fmt.Println(version, commit, date)
	}
	conf := internal_moonmist.GetGenConf(viper.GetString(fileFlag))
	internal_moonmist.MkdirAll(conf)
	currentGoModulePath := parser.GetModulePath()
	allPackages := make(map[string]*parser.Package)
	runtime.Must(filepath.Walk(conf.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			return nil
		}
		if conf.Dir == path {
			return nil
		}
		pkg, err := parser.ParsePackage(path, currentGoModulePath)
		if err != nil {
			panic(err)
		}
		allPackages[pkg.Name] = pkg
		return nil
	}))
	for _, pkg := range allPackages {
		for _, c := range conf.Codes {
			switch c.Type {
			case "ts":
				runtime.Must(parser.GenerateTSCode(conf.Dir, pkg, allPackages))
			case "dart":
				runtime.Must(parser.GenerateDartCode(conf.Dir, pkg, allPackages))
			case "swift":
				runtime.Must(parser.GenerateSwiftCode(conf.Dir, pkg, allPackages))
			case "chi_route":
				runtime.Must(parser.GenerateChiRoutes(conf.Dir, pkg, allPackages))
			case "angular_delon":
				runtime.Must(parser.GenerateTSAngularDelonCode(conf.Dir, pkg, allPackages))
			case "proto":
				runtime.Must(parser.GenerateProtoFile(conf.Dir, currentGoModulePath, pkg, allPackages))
			case "gsrv":
				runtime.Must(parser.GenerateGsrv(conf.Dir, currentGoModulePath, pkg, allPackages))
			default:
				panic("unsupported " + c.Type)
			}
		}
	}
}
