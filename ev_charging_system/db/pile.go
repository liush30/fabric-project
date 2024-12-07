package db

// Pile 充电桩信息表
type Pile struct {
	PileID      string `gorm:"primaryKey;type:varchar(64);not null"`
	StationID   string `gorm:"type:varchar(20)"`
	PileCode    string `gorm:"type:varchar(64);not null"`
	PileName    string `gorm:"type:varchar(30);not null"`
	Description string `gorm:"type:varchar(100)"`
	Location    string `gorm:"type:varchar(100);not null"`
	Status      string `gorm:"type:varchar(10);not null"`
	Type        string `gorm:"type:varchar(64);not null"`
}

func (Pile) TableName() string {
	return "t_pile"
}
