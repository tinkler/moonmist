package parser

import (
	"os/exec"
	"strings"
)

func GetModulePath(root string) string {
	cmd := exec.Command("go", "list", "-m")
	cmd.Dir = root
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	moudles := strings.Split(string(output), "\n")
	root = strings.ReplaceAll(root, "\\", "/")
	for _, m := range moudles {
		if strings.HasSuffix(root, m) {
			return m
		}
	}
	return strings.TrimSpace(string(output))
}
