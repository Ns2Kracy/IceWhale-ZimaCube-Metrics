package service

type Metrics struct{}

func NewMetrics() *Metrics {
	return &Metrics{}
}

func (h *Metrics) GetZimaCubeIP() string {
	return ""
}
