package service

import (
	"github.com/IceWhaleTech/CasaOS-Common/external"
)

var (
	Gateway         external.ManagementService
	ZimaCubeMetrics *Metrics
)

func Initialize(RuntimePath string) {
	_gateway, err := external.NewManagementService(RuntimePath)
	if err != nil && len(RuntimePath) > 0 {
		panic(err)
	}

	Gateway = _gateway

	ZimaCubeMetrics = NewMetrics()
}
