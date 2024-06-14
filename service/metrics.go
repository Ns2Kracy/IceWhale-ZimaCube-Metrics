package service

import (
	"fmt"
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

func (m *Metrics) GetMetric(serviceName string) codegen.Metric {
	var metrics model.MetricDBModel
	m.DB.Where("name = ?", serviceName).Last(&metrics)

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
	m.DB.Where("name = ? AND created_at > ?", serviceName, time.Now().Add(-1*time.Hour)).Order("cpu desc").First(&metrics)

	return metrics.CPU + "%"
}

func (m *Metrics) GetAvgCPU(serviceName string) string {
	var metrics model.MetricDBModel
	// Get the average cpu from the database
	m.DB.Where("name = ?", serviceName).Select("AVG(cpu)").Find(&metrics)

	return metrics.CPU + "%"
}

func (m *Metrics) GetMaxMem(serviceName string) string {
	var metrics model.MetricDBModel
	// Get the max cpu and mem from the database in the last 5 minutes
	m.DB.Where("name = ? AND created_at > ?", serviceName, time.Now().Add(-5*time.Minute)).Order("mem desc").First(&metrics)

	return metrics.MEM + " MB"
}

func (m *Metrics) GetAvgMem(serviceName string) string {
	var metrics model.MetricDBModel
	// Get the average mem from the database
	m.DB.Select("AVG(mem)").Where("name = ?", serviceName).Find(&metrics)

	return metrics.MEM + " MB"
}

func (m *Metrics) GetMetrics() []codegen.Metric {
	metrics := []codegen.Metric{}
	for _, service := range common.ServiceList {
		metric := m.GetMetric(service)
		metrics = append(metrics, metric)
	}

	return metrics
}
