package model

type Parameter struct {
	ParamID    string `gorm:"primaryKey;type:varchar(64);not null"`
	ParamKey   string `gorm:"type:varchar(64);not null"`
	ParamValue string `gorm:"type:varchar(10);not null"`
}

func (Parameter) TableName() string {
	return "t_parameter"
}
