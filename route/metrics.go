package route

import "github.com/labstack/echo/v4"

func (m *Metrics) GetMetrics(c echo.Context) error {
	return c.JSON(200, "Metrics")
}
