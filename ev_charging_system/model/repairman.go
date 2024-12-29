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
	RepairmanID string    `gorm:"primaryKey;type:varchar(64);not null" json:"repairman_id"`
	UserName    string    `gorm:"type:varchar(60);not null" json:"user_name"`
	Password    string    `gorm:"type:varchar(255);not null;comment:'密码'" json:"password"`
	Name        string    `gorm:"type:varchar(20);default:null" json:"name"`
	ContactInfo string    `gorm:"type:varchar(64);not null" json:"contact_info"`
	Status      int8      `gorm:"type:tinyint;not null;default:0;comment:'0，空闲中、1工作中、2已休息、3已离职'" json:"status"`
	Description string    `gorm:"type:varchar(200);default:null" json:"description"`
	UserType    int8      `gorm:"type:tinyint;not null;default:0;comment:'用户类型，0，维修员，1，充电站 2、管理员'" json:"user_type"`
	RegistTime  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null" json:"regist_time"`
}

func (Repairman) TableName() string {
	return "t_repairman"
}
