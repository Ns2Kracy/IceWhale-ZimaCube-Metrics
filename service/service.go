package service

import (
	"github.com/IceWhaleTech/CasaOS-Common/external"
	"gorm.io/gorm"
)

var MyService *Services

type Services struct {
	gateway     external.ManagementService
	metrics     *Metrics
	db          *gorm.DB
	runtimePath string
}

func Initialize(runtimePath string) {
	MyService = &Services{
		runtimePath: runtimePath,
	}
}

func (s *Services) Gateway() external.ManagementService {
	if s.gateway == nil {
		gateway, err := external.NewManagementService(s.runtimePath)
		if err != nil && len(s.runtimePath) > 0 {
			panic(err)
		}

		s.gateway = gateway
	}

	return s.gateway
}

func (s *Services) Metrics() *Metrics {
	if s.metrics == nil {
		s.metrics = NewMetrics(s.db)
	}

	return s.metrics
}
