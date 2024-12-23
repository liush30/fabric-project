package model

type FeeRule struct {
	RuleID       string  `gorm:"primaryKey;type:varchar(64);not null"`
	StationID    string  `gorm:"type:varchar(64);not null"`
	RuleName     string  `gorm:"type:varchar(20);not null"`
	ChargingType string  `gorm:"type:varchar(10);not null"`
	UnitPrice    string  `gorm:"type:varchar(50)"`
	PeakRate     float64 `gorm:"type:decimal(10,2)"`
	OffPeakRate  float64 `gorm:"type:decimal(10,2)"`
	MinCharge    float64 `gorm:"type:decimal(10,2)"`
	MaxCharge    float64 `gorm:"type:decimal(10,2)"`
	StartTime    string  `gorm:"type:varchar(26)"`
	EndTime      string  `gorm:"type:varchar(26)"`
}

func (FeeRule) TableName() string {
	return "t_fee_rule"
}
