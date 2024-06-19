package utils

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

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

func GetProcessInfo(processPid string) (string, string, string) {
	pid, _ := strconv.ParseInt(processPid, 10, 32)
	p, err := process.NewProcess(int32(pid))
	if err != nil {
		fmt.Println("Error getting process info:", err)
		return "", "", ""
	}

	cpu, err := p.CPUPercent()
	if err != nil {
		fmt.Println("Error getting CPU percent:", err)
		return "", "", ""
	}

	memoryInfo, err := p.MemoryInfo()
	memUsage := memoryInfo.RSS

	if err != nil {
		fmt.Println("Error getting memory percent:", err)
		return "", "", ""
	}

	createTime, err := p.CreateTime()
	if err != nil {
		return "", "", ""
	}
	createTimeUnix := time.Unix(0, createTime*int64(time.Microsecond))
	uptime := time.Since(createTimeUnix)
	uptimeH := time.Duration(uptime.Hours())
	uptimeM := time.Duration(uptime.Minutes())
	uptimeS := time.Duration(uptime.Seconds())

	uptimeString := fmt.Sprintf("%02dh %02dm %02ds", int(uptimeH), int(uptimeM)%60, int(uptimeS)%60)

	return fmt.Sprintf("%.2f", cpu), fmt.Sprintf("%.2f", float64(memUsage)/1024/1024), uptimeString
}
