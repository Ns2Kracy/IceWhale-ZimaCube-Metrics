package model

import "time"

type MetricDBModel struct {
	Name      string    `json:"name" gorm:"string"`
	CPU       string    `json:"cpu" gorm:"string"`
	MEM       string    `json:"mem" gorm:"string"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"autoCreateTime"`
}

func (s *MetricDBModel) TableName() string {
	return "o_metrics"
}
