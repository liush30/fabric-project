package db

// Repairman 维修员信息表
type Repairman struct {
	RepairmanID string `gorm:"primaryKey;type:varchar(64);not null"`
	Name        string `gorm:"type:varchar(20)"`
	ContactInfo string `gorm:"type:varchar(64);not null"`
	Status      string `gorm:"type:varchar(10);not null"`
	Description string `gorm:"type:varchar(200)"`
	RegistTime  string `gorm:"type:varchar(26);not null"`
}

func (Repairman) TableName() string {
	return "t_repairman"
}
