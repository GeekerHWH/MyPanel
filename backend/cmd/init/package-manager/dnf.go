package packagemanager

import (
	"os/exec"
	"strings"
)

func checkDNFInstalled() bool {
	cmd := exec.Command("rpm", "-q", "dnf")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	if strings.Contains(string(output), "is installed") {
		return true
	}
	return false
}
