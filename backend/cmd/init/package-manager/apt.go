package packagemanager

import (
	"os/exec"
	"strings"
)

func checkAPTInstalled() bool {
	cmd := exec.Command("dpkg", "-s", "apt")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	if strings.Contains(string(output), "Status: install ok installed") {
		return true
	}
	return false
}
