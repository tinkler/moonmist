package parser

import (
	"os"
	"testing"
)

func TestGetModulePath(t *testing.T) {
	if err := os.Chdir("../.."); err != nil {
		t.Fatal(err)
	}
	root, _ := os.Getwd()
	if GetGoModule(root) != "github.com/tinkler/mqttadmin" {
		t.Fail()
	}
}
