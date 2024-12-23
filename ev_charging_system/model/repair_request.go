package model

type RepairRequest struct {
	RepairID    string `gorm:"primaryKey;type:varchar(64);not null"`
	StationID   string `gorm:"type:varchar(64);not null"`
	PileID      string `gorm:"type:varchar(64);not null"`
	RepairmanID string `gorm:"type:varchar(64);not null"`
	Description string `gorm:"type:varchar(200)"`
	Status      string `gorm:"type:varchar(10);not null"`
	Reason      string `gorm:"type:varchar(200)"`
	Result      string `gorm:"type:varchar(200)"`
	RequestTime string `gorm:"type:varchar(26);not null"`
	EndTime     string `gorm:"type:varchar(26)"`
}

func (RepairRequest) TableName() string {
	return "t_repair_request"
}
