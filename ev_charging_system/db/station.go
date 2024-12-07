package db

// Station 充电站基础信息表
type Station struct {
	StationID     string `gorm:"primaryKey;type:varchar(64);not null"`
	StationName   string `gorm:"type:varchar(20);not null"`
	Location      string `gorm:"type:varchar(100);not null"`
	City          string `gorm:"type:varchar(20);not null"`
	District      string `gorm:"type:varchar(20);not null"`
	ContactNumber string `gorm:"type:varchar(26);not null"`
	ManagerName   string `gorm:"type:varchar(10);not null"`
	OpeningHours  string `gorm:"type:varchar(20);not null"`
	Status        string `gorm:"type:varchar(10)"`
	Description   string `gorm:"type:varchar(100);not null"`
	LoginPwd      string `gorm:"type:varchar(30);not null"`
}

func (Station) TableName() string {
	return "t_station"
}
