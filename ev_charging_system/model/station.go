package model

// Station 充电站基础信息表
//type Station struct {
//	StationID     string `gorm:"primaryKey;type:varchar(64);not null"`
//	StationName   string `gorm:"type:varchar(20);not null"`
//	Location      string `gorm:"type:varchar(100);not null"`
//	City          string `gorm:"type:varchar(20);not null"`
//	District      string `gorm:"type:varchar(20);not null"`
//	ContactNumber string `gorm:"type:varchar(26);not null"`
//	ManagerName   string `gorm:"type:varchar(10);not null"`
//	OpeningHours  string `gorm:"type:varchar(20);not null"`
//	Status        string `gorm:"type:varchar(10)"`
//	Description   string `gorm:"type:varchar(100);not null"`
//	LoginPwd      string `gorm:"type:varchar(30);not null"`
//}

type Station struct {
	StationID     string `gorm:"column:station_id;primaryKey;size:64" json:"stationId"`
	RepairmanID   string `gorm:"column:t_repairman;size:64" json:"repairmanId,omitempty"` // 负责人id
	StationName   string `gorm:"column:station_name;size:20;not null" json:"stationName"`
	Location      string `gorm:"column:location;size:100;not null" json:"location"`
	City          string `gorm:"column:city;size:20;not null" json:"city"`
	District      string `gorm:"column:district;size:20;not null" json:"district"`
	ContactNumber string `gorm:"column:contact_number;size:26;not null" json:"contactNumber"`
	ManagerName   string `gorm:"column:manager_name;size:10;not null" json:"managerName"`
	OpeningHours  string `gorm:"column:opening_hours;size:20;not null" json:"openingHours"`
	Status        int8   `gorm:"column:status;default:0" json:"status"` // 0 关闭、1 暂停营业、2 运营中
	Description   string `gorm:"column:description;size:100;not null" json:"description"`
	LoginPwd      string `gorm:"column:login_pwd;size:30;not null" json:"loginPwd"`
}

func (Station) TableName() string {
	return "t_station"
}
