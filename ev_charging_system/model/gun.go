package model

type Gun struct {
	GunID      string  `gorm:"primaryKey;type:varchar(64);not null"`
	PileID     string  `gorm:"type:varchar(64);not null"`
	GunType    string  `gorm:"type:varchar(64);not null"`
	GunStatus  string  `gorm:"type:varchar(30);not null"`
	Power      string  `gorm:"type:varchar(50)"`
	Voltage    float64 `gorm:"type:decimal(10,2)"`
	Current    float64 `gorm:"type:decimal(10,2)"`
	MaxCurrent float64 `gorm:"type:decimal(10,2)"`
}

func (Gun) TableName() string {
	return "t_gun"
}
