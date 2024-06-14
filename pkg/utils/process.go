package utils

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/v4/process"
)

func GetPid(service string) string {
	cmd := exec.Command("pidof", service)
	output, err := cmd.Output()
	if err != nil {
		return "-1"
	}

	return strings.TrimSpace(string(output))
}

func GetProcessInfo(processPid string) (string, string) {
	pid, _ := strconv.ParseInt(processPid, 10, 32)
	p, err := process.NewProcess(int32(pid))
	if err != nil {
		fmt.Println("Error getting process info:", err)
		return "", ""
	}

	cpu, err := p.CPUPercent()
	if err != nil {
		fmt.Println("Error getting CPU percent:", err)
		return "", ""
	}

	memoryInfo, err := p.MemoryInfo()
	memUsage := memoryInfo.RSS

	if err != nil {
		fmt.Println("Error getting memory percent:", err)
		return "", ""
	}

	// Convert to MB
	return fmt.Sprintf("%.2f", cpu), fmt.Sprintf("%.2f", float64(memUsage)/1024/1024)
}
