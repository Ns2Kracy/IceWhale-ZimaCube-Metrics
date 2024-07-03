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

func GetProcessInfo(processPid string) map[string]string {
	pi := make(map[string]string)

	pid, _ := strconv.ParseInt(processPid, 10, 32)
	p, err := process.NewProcess(int32(pid))
	if err != nil {
		fmt.Println("Error getting process info:", err)
		return nil
	}

	// Get CPU percent
	cpu, err := p.CPUPercent()
	if err != nil {
		fmt.Println("Error getting CPU percent:", err)
		return nil
	}

	// Get memory usage
	memoryInfo, err := p.MemoryInfo()
	memUsage := memoryInfo.RSS
	if err != nil {
		fmt.Println("Error getting memory percent:", err)
		return nil
	}

	createTime, err := p.CreateTime()
	if err != nil {
		return nil
	}

	createTimeUnix := time.Unix(0, createTime*int64(time.Millisecond))
	uptime := time.Since(createTimeUnix)
	uptimeH := time.Duration(uptime.Hours())
	uptimeM := time.Duration(uptime.Minutes())
	uptimeS := time.Duration(uptime.Seconds())
	uptimeString := fmt.Sprintf("%02dh %02dm %02ds", int(uptimeH), int(uptimeM)%60, int(uptimeS)%60)

	// Get Temprature

	pi["cpu"] = fmt.Sprintf("%.2f", cpu)
	pi["mem"] = fmt.Sprintf("%.2f", float64(memUsage)/1024/1024)
	pi["uptime"] = uptimeString

	return pi
}
