package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/codegen"
	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/common"
	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/pkg/utils"
	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/service/model"
	"gorm.io/gorm"
)

type Metrics struct {
	DB *gorm.DB
}

func NewMetrics(db *gorm.DB) *Metrics {
	return &Metrics{
		DB: db,
	}
}

func (m *Metrics) Monitor() {
	for {
		for _, service := range common.ServiceList {
			pid := utils.GetPid(service)
			if pid == "-1" {
				fmt.Println("Service", service, "is not running.")
				continue
			}

			cpu, mem := utils.GetProcessInfo(pid)

			m.DB.Create(&model.MetricDBModel{
				Name: service,
				CPU:  cpu,
				MEM:  mem,
			})
		}
		time.Sleep(1 * time.Second)
	}
}

func (m *Metrics) ReportFeiShu(webhookURL string) {
	// for {
	// 	for _, service := range common.ServiceList {
	// 		metrics := m.GetMetric(service)
	// 		count := 0

	// 		cpu, _ := strconv.ParseFloat(*metrics.Cpu, 64)
	// 		mem, _ := strconv.ParseFloat(*metrics.Mem, 64)
	// 		maxCPU, _ := strconv.ParseFloat(*metrics.MaxCpu, 64)
	// 		maxMem, _ := strconv.ParseFloat(*metrics.MaxMem, 64)

	// 		if cpu > maxCPU*0.5 || mem > maxMem*0.5 {
	// 			count++
	// 			if count == 60 {
	// 				message := strings.Join([]string{
	// 					"Service: ", service, " CPU/MEM 超过阈值",
	// 					"CPU: ", *metrics.Cpu, " MEM: ", *metrics.Mem,
	// 				}, "")
	// 				_ = utils.SendTextMessage(webhookURL, message)
	// 				count = 0
	// 			}
	// 		}
	// 	}
	// 	time.Sleep(5 * time.Second)
	// }
}

func (m *Metrics) GetMetric(serviceName string) codegen.Metric {
	var metrics model.MetricDBModel
	m.DB.Where("name = ?", serviceName).Order("created_at desc").Last(&metrics)

	name := metrics.Name
	cpu := metrics.CPU + "%"
	mem := metrics.MEM + " MB"
	avgCPU := m.GetAvgCPU(serviceName)
	maxCPI := m.GetMaxCPU(serviceName)
	avgMem := m.GetAvgMem(serviceName)
	maxMem := m.GetMaxMem(serviceName)

	return codegen.Metric{
		Name:   &name,
		Cpu:    &cpu,
		Mem:    &mem,
		AvgCpu: &avgCPU,
		MaxCpu: &maxCPI,
		AvgMem: &avgMem,
		MaxMem: &maxMem,
	}
}

func (m *Metrics) GetMaxCPU(serviceName string) string {
	var metrics model.MetricDBModel
	// Get the max cpu and mem from the database in the last 1 hour
	m.DB.Where("name = ?", serviceName).Order("cpu desc").First(&metrics)

	maxCPU, _ := strconv.ParseFloat(metrics.CPU, 64)

	return fmt.Sprintf("%.2f", maxCPU) + "%"
}

func (m *Metrics) GetAvgCPU(serviceName string) string {
	var metrics model.MetricDBModel
	// Get the average cpu from the database
	m.DB.Raw("SELECT AVG(cpu) as cpu FROM o_metrics WHERE name = ?", serviceName).Scan(&metrics)

	avgCPU, _ := strconv.ParseFloat(metrics.CPU, 64)

	return fmt.Sprintf("%.2f", avgCPU) + "%"
}

func (m *Metrics) GetMaxMem(serviceName string) string {
	var metrics model.MetricDBModel
	// Get the max mem from the database
	m.DB.Where("name = ?", serviceName).Order("mem desc").First(&metrics)

	maxMem, _ := strconv.ParseFloat(metrics.MEM, 64)

	return fmt.Sprintf("%.2f", maxMem) + " MB"
}

func (m *Metrics) GetAvgMem(serviceName string) string {
	var metrics model.MetricDBModel
	// Get the average mem from the database
	m.DB.Raw("SELECT AVG(mem) as mem FROM o_metrics WHERE name = ?", serviceName).Scan(&metrics)

	avgMem, _ := strconv.ParseFloat(metrics.MEM, 64)

	return fmt.Sprintf("%.2f", avgMem) + " MB"
}

func (m *Metrics) GetMetrics() []codegen.Metric {
	metrics := []codegen.Metric{}
	for _, service := range common.ServiceList {
		metric := m.GetMetric(service)
		metrics = append(metrics, metric)
	}

	return metrics
}

func (m *Metrics) AddSshClient(params codegen.AddZimaCube) {
}
