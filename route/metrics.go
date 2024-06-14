package route

import (
	"net/http"

	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/codegen"
	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/service"
	"github.com/labstack/echo/v4"
)

func (m *MetricsRoute) GetMetrics(c echo.Context) error {
	metrics := service.MyService.Metrics().GetMetrics()

	return c.JSON(http.StatusOK, codegen.ResponseZimaCubeMetricsOK{
		Data: &metrics,
	})
}
