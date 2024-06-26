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
	flagVersion = "version"
	flagFile    = "file"
)

func parseEnvFlags() {
	pflag.Bool(flagVersion, false, "print version and exit")
	pflag.StringP(flagFile, "f", "mist.yaml", "moonmist config file")
	pflag.Parse()

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	runtime.Must(viper.BindPFlags(pflag.CommandLine))

}

func main() {
	parseEnvFlags()
	if viper.GetBool(flagVersion) {
		fmt.Println(version, commit, date)
		return
	}
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if len(os.Args) > 1 {
		root = filepath.Join(root, os.Args[1])
	}
	conf := internal_moonmist.GetGenConf(viper.GetString(flagFile), root)
	internal_moonmist.MkdirAll(conf)
	currentGoModule := parser.GetGoModule(root)
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
		pkg, err := parser.ParsePackage(path, currentGoModule)
		if err != nil {
			panic(err)
		}
		if pkg.Name == "" {
			return nil
		}
		allPackages[pkg.Name] = pkg
		return nil
	}))
	for _, pkg := range allPackages {
		for _, c := range conf.Codes {
			switch c.Type {
			case "ts":
				runtime.Must(parser.GenerateTSCode(c.Out, pkg, allPackages))
			case "dart":
				runtime.Must(parser.GenerateDartCode(c.Out, pkg, allPackages))
			case "swift":
				runtime.Must(parser.GenerateSwiftCode(c.Out, pkg, allPackages))
			case "chi_route":
				runtime.Must(parser.GenerateChiRoutes(c.Out, pkg, allPackages))
			case "angular_delon":
				runtime.Must(parser.GenerateTSAngularDelonCode(c.Out, pkg, allPackages))
			case "proto":
				runtime.Must(parser.GenerateProtoFile(c.Out, root, currentGoModule, pkg, allPackages))
			case "gsrv":
				runtime.Must(parser.GenerateGsrv(c.Out, currentGoModule, pkg, allPackages))
			default:
				panic("unsupported " + c.Type)
			}
		}
	}
	fmt.Println("done")
}
