package model

import "time"

// Repairman 维修员信息表
//type Repairman struct {
//	RepairmanID string `gorm:"primaryKey;type:varchar(64);not null"`
//	Name        string `gorm:"type:varchar(20)"`
//	ContactInfo string `gorm:"type:varchar(64);not null"`
//	Status      string `gorm:"type:varchar(10);not null"`
//	Description string `gorm:"type:varchar(200)"`
//	RegistTime  string `gorm:"type:varchar(26);not null"`
//}

type Repairman struct {
	RepairmanID string    `gorm:"column:repairman_id;primaryKey;size:64" json:"repairmanId"`
	UserName    string    `gorm:"column:user_name;size:60;not null" json:"userName"`
	Password    string    `gorm:"column:password;size:255;not null" json:"password"` // 密码
	Name        string    `gorm:"column:name;size:20" json:"name,omitempty"`
	ContactInfo string    `gorm:"column:contact_info;size:64;not null" json:"contactInfo"`
	Status      int8      `gorm:"column:status;not null;default:0" json:"status"` // 0 空闲中、1 工作中、2 已休息、3 已离职
	Description string    `gorm:"column:description;size:200" json:"description,omitempty"`
	UserType    int8      `gorm:"column:user_type;not null;default:0" json:"userType"`  // 用户类型：0 维修员、1 充电站、2 管理员
	StationID   string    `gorm:"column:station_Id;size:64" json:"stationId,omitempty"` // 所属充电站id
	RegistTime  time.Time `gorm:"column:regist_time;not null;default:CURRENT_TIMESTAMP" json:"registTime"`
}

func (Repairman) TableName() string {
	return "t_repairman"
}
