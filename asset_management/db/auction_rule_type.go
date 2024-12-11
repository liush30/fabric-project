package db

import "fmt"

type AuctionRuleType struct {
	TypeID      string `gorm:"primaryKey;column:type_id"` // 类型ID
	TypeStatus  string `gorm:"column:type_status"`        // 类型状态
	TypeName    string `gorm:"column:type_name"`          // 类型名称
	TypeContent string `gorm:"column:type_content"`       // 类型说明
}

func (AuctionRuleType) TableName() string {
	return "auction_rule_type"
}

func CreateAuctionRuleType(AuctionRuleType *AuctionRuleType) error {
	result := DB.Create(AuctionRuleType)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("insert failed,affected rows:%d", result.RowsAffected)
	}
	return nil
}

// GetAuctionRuleTypeByID 查询拍卖规则类型
func GetAuctionRuleTypeByID(id string) (*AuctionRuleType, error) {
	var AuctionRuleType AuctionRuleType
	err := DB.Model(&AuctionRuleType).First(&AuctionRuleType, "type_id = ?", id).Error
	return &AuctionRuleType, err
}

// UpdateAuctionRuleType 更新拍卖规则类型
func UpdateAuctionRuleType(id string, updatedFields interface{}) error {
	return DB.Model(&AuctionRuleType{}).Where("type_id = ?", id).Updates(updatedFields).Error
}

// DeleteAuctionRuleType 删除拍卖规则类型
func DeleteAuctionRuleType(id string) error {
	return DB.Delete(&AuctionRuleType{}, "type_id = ?", id).Error
}
func GetAllAuctionRuleTypeWithConditions(conditions map[string]interface{}, page, pageSize int) ([]AuctionRuleType, error) {
	var task []AuctionRuleType

	// 计算 OFFSET
	offset := (page - 1) * pageSize

	// 使用 Limit 和 Offset 进行分页查询，并根据条件构建 WHERE 子句
	err := DB.Model(&task).Where(conditions).Limit(pageSize).Offset(offset).Find(&task).Error
	return task, err
}
