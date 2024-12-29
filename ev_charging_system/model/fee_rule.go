package model

type FeeRule struct {
	RuleID       string  `gorm:"column:rule_id;primaryKey;size:64" json:"ruleId"`
	StationID    string  `gorm:"column:station_id;size:64;not null" json:"stationId"`
	RuleName     string  `gorm:"column:rule_name;size:20;not null" json:"ruleName"`
	ChargingType int8    `gorm:"column:charging_type;default:0;not null" json:"chargingType"` // 时间计费、按电量记录、混合计费
	UnitPrice    string  `gorm:"column:unit_price;size:50" json:"unitPrice,omitempty"`
	PeakRate     float64 `gorm:"column:peak_rate;type:decimal(10,2)" json:"peakRate,omitempty"`
	OffPeakRate  float64 `gorm:"column:off_peak_rate;type:decimal(10,2)" json:"offPeakRate,omitempty"`
	MinCharge    float64 `gorm:"column:min_charge;type:decimal(10,2)" json:"minCharge,omitempty"`
	MaxCharge    float64 `gorm:"column:max_charge;type:decimal(10,2)" json:"maxCharge,omitempty"`
	StartTime    int64   `gorm:"column:start_time" json:"startTime,omitempty"`
	EndTime      int64   `gorm:"column:end_time" json:"endTime,omitempty"`
}

func (FeeRule) TableName() string {
	return "t_fee_rule"
}
