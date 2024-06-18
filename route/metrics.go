package route

import (
	"net/http"

	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/codegen"
	"github.com/Ns2Kracy/IceWhale-ZimaCube-Metrics/service"
	"github.com/labstack/echo/v4"
)

func (m *MetricsRoute) GetMetrics(ctx echo.Context) error {
	metrics := service.MyService.Metrics().GetMetrics()

	return ctx.JSON(http.StatusOK, codegen.ResponseZimaCubeMetricsOK{
		Data: &metrics,
	})
}

func (m *MetricsRoute) PostAddZimaCube(ctx echo.Context) error {
	return ctx.NoContent(http.StatusCreated)
}

func (m *MetricsRoute) DeleteZimaCube(ctx echo.Context, params codegen.DeleteZimaCubeParams) error {
	return ctx.NoContent(http.StatusNoContent)
}
