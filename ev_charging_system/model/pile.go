package model

// Pile 充电桩信息表
type Pile struct {
	PileID      string `gorm:"primaryKey;type:varchar(64);not null"`
	StationID   string `gorm:"type:varchar(64)"`
	PileCode    string `gorm:"type:varchar(64);not null"`
	PileName    string `gorm:"type:varchar(30);not null"`
	Description string `gorm:"type:varchar(100)"`
	Location    string `gorm:"type:varchar(100);not null"`
	Status      int8   `gorm:"column:status;default:0" json:"status"`
	Type        int8   `gorm:"column:type;default:0" json:"type"`
}

func (Pile) TableName() string {
	return "t_pile"
}
