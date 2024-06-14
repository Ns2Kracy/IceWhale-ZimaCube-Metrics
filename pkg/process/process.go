package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetProcessInfo(pid string) (string, string) {
	cmd := exec.Command("ps", "-p", pid, "-o", "pcpu=,pmem=")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running ps command:", err)
		return "", ""
	}

	out := string(output)
	lines := strings.Split(out, "\n")
	fields := strings.Fields(lines[0])

	return fields[0], fields[1]
}
