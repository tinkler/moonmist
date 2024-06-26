package parser

import (
	"os"
	"testing"
)

func TestParsePackage(t *testing.T) {
	if err := os.Chdir("../.."); err != nil {
		t.Fatal(err)
	}
	root, _ := os.Getwd()
	pkg, err := ParsePackage("pkg/model/user", GetGoModule(root))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pkg)
}

func TestParsePackageWithStream(t *testing.T) {
	if err := os.Chdir("../.."); err != nil {
		t.Fatal(err)
	}
	root, _ := os.Getwd()
	pkg, err := ParsePackage("pkg/model/page", GetGoModule(root))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pkg)
}

func TestParseMapType(t *testing.T) {
	mapStr := "map[string]Chapter"
	match := mapTypeRe.FindStringSubmatch(mapStr)
	if match[1] != "string" {
		t.Fail()
	}
	if match[2] != "Chapter" {
		t.Fail()
	}
}
