package model

type MetricDBModel struct {
	Name      string `json:"name" gorm:"string"`
	CPU       string `json:"cpu" gorm:"string"`
	MEM       string `json:"mem" gorm:"string"`
	CreatedAt string `json:"created_at,omitempty" gorm:"<-:create;autoCreateTime"`
}

func (s *MetricDBModel) TableName() string {
	return "o_metrics"
}
