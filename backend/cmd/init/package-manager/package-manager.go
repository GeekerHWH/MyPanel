package packagemanager

import (
	"fmt"
	"os/exec"
	"strings"
)

func Init() {
	fmt.Println(isDPKGInstalled())
	fmt.Println(isRPMInstalled())
	fmt.Println(checkAPTInstalled())
	fmt.Println(checkDNFInstalled())
}

func isDPKGInstalled() bool {
	cmd := exec.Command("which", "dpkg")
	return cmd.Run() == nil
}

func isRPMInstalled() bool {
	cmd := exec.Command("which", "rpm")
	return cmd.Run() == nil
}

func checkInstalledByDPKG(packageName string) bool {
	cmd := exec.Command("dpkg", "-s", packageName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}
	if strings.Contains(string(output), "Status: install ok installed") {
		return true
	}
	return false
}
