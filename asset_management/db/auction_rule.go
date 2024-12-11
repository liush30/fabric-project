package db

import "fmt"

type AuctionRule struct {
	RuleID       string `gorm:"primaryKey;column:rule_id"` // 规则ID
	RuleTitle    string `gorm:"column:rule_title"`         // 规则标题
	RuleContent  string `gorm:"column:rule_content"`       // 规则说明
	BidIncrement string `gorm:"column:bid_increment"`      // 最小加价增量
	MaxBidCount  string `gorm:"column:max_bid_count"`      // 最大出价次数
	StartPrice   string `gorm:"column:start_price"`        // 起拍价
	TimeLimit    string `gorm:"column:time_limit"`         // 拍卖时限
	AutoExtend   string `gorm:"column:auto_extend"`        // 是否启用自动延时
	TypeID       string `gorm:"column:type_id"`            // 拍卖类型ID
	EnableAsset  string `gorm:"column:enable_asset"`       // 是否启用资产限制
	AssetLimit   int    `gorm:"column:asset_limit"`        // 资产限制值
	EnableCredit string `gorm:"column:enable_credit"`      // 是否启用信用评分限制
	ScoreLimit   int    `gorm:"column:score_limit"`        // 信用评分限制值
}

func (AuctionRule) TableName() string {
	return "auction_rule"
}
func CreateAuctionRule(AuctionRule *AuctionRule) error {
	result := DB.Create(AuctionRule)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("insert failed,affected rows:%d", result.RowsAffected)
	}
	return nil
}

// GetAuctionRuleByID 查询拍卖规则
func GetAuctionRuleByID(id string) (*AuctionRule, error) {
	var AuctionRule AuctionRule
	err := DB.Model(&AuctionRule).First(&AuctionRule, "rule_id = ?", id).Error
	return &AuctionRule, err
}

// UpdateAuctionRule 更新拍卖规则
func UpdateAuctionRule(id string, updatedFields interface{}) error {
	return DB.Model(&AuctionRule{}).Where("rule_id = ?", id).Updates(updatedFields).Error
}

// DeleteAuctionRule 删除拍卖规则
func DeleteAuctionRule(id string) error {
	return DB.Delete(&AuctionRule{}, "rule_id = ?", id).Error
}
func GetAllAuctionRuleWithConditions(conditions map[string]interface{}, page, pageSize int) ([]AuctionRule, error) {
	var task []AuctionRule

	// 计算 OFFSET
	offset := (page - 1) * pageSize

	// 使用 Limit 和 Offset 进行分页查询，并根据条件构建 WHERE 子句
	err := DB.Model(&task).Where(conditions).Limit(pageSize).Offset(offset).Find(&task).Error
	return task, err
}
