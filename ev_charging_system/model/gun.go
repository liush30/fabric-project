package model

type Gun struct {
	GunID      string  `gorm:"column:gun_id;primaryKey;size:64" json:"gunId"`
	PileID     string  `gorm:"column:pile_id;size:64;not null" json:"pileId"`
	GunType    int8    `gorm:"column:gun_type;default:1" json:"gunType"`     // 1 正常、2 故障、3 维修中
	GunStatus  int8    `gorm:"column:gun_status;default:0" json:"gunStatus"` // AC 和 DC
	Power      string  `gorm:"column:power;size:50" json:"power,omitempty"`
	Voltage    float64 `gorm:"column:voltage;type:decimal(10,2)" json:"voltage,omitempty"`
	Current    float64 `gorm:"column:current;type:decimal(10,2)" json:"current,omitempty"`
	MaxCurrent float64 `gorm:"column:max_current;type:decimal(10,2)" json:"maxCurrent,omitempty"`
}

func (Gun) TableName() string {
	return "t_gun"
}
