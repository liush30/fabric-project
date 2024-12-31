package model

type RepairRequest struct {
	RepairID    string `gorm:"column:repair_id;primaryKey;size:64" json:"repairId"`
	StationID   string `gorm:"column:station_id;size:64;not null" json:"stationId"`
	PileID      string `gorm:"column:pile_id;size:64;not null" json:"pileId"`
	RepairmanID string `gorm:"column:repairman_id;size:64;not null" json:"repairmanId"`
	Description string `gorm:"column:description;size:200" json:"description,omitempty"`
	Status      int8   `gorm:"column:status;default:0" json:"status"` // 0待处理、1已取消、2处理中、3已完成
	Reason      string `gorm:"column:reason;size:200" json:"reason,omitempty"`
	Result      string `gorm:"column:result;size:200" json:"result,omitempty"`
	RequestTime string `gorm:"column:request_time" json:"requestTime,omitempty"`
	EndTime     string `gorm:"column:end_time" json:"endTime,omitempty"`
}

func (RepairRequest) TableName() string {
	return "t_repair_request"
}
