package db

import "fmt"

type BuyerRequest struct {
	RequestID        string `gorm:"primaryKey;column:request_id"` // 请求ID
	UserID           string `gorm:"column:user_id"`               // 用户ID
	BuyerName        string `gorm:"column:buyer_name"`            // 竞买人姓名
	BuyerPhone       string `gorm:"column:buyer_phone"`           // 竞买人联系方式
	BuyerNumber      string `gorm:"column:buyer_number"`          // 竞买人身份证号码
	BuyerAddress     string `gorm:"column:buyer_address"`         // 竞买人地址
	RejectionContent string `gorm:"column:rejection_content"`     // 被拒原因
	TotalAssets      int    `gorm:"column:total_assets"`          // 总资产
	CreditScore      int    `gorm:"column:credit_score"`          // 信用评分
	IncomeSource     string `gorm:"column:income_source"`         // 收入来源
	Result           string `gorm:"column:result"`                // 认证结果
	RequestTime      string `gorm:"column:request_time"`          // 申请时间
	ProcessTime      string `gorm:"column:process_time"`          // 处理时间
}

func (BuyerRequest) TableName() string {
	return "buyer_request"
}

func CreateBuyerRequest(BuyerRequest *BuyerRequest) error {
	result := DB.Create(BuyerRequest)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("insert failed,affected rows:%d", result.RowsAffected)
	}
	return nil
}

// GetBuyerRequestByID 查询竞买人申请信息
func GetBuyerRequestByID(id string) (*BuyerRequest, error) {
	var BuyerRequest BuyerRequest
	err := DB.Model(&BuyerRequest).First(&BuyerRequest, "request_id = ?", id).Error
	return &BuyerRequest, err
}

// UpdateBuyerRequest 更新竞买人申请信息
func UpdateBuyerRequest(id string, updatedFields interface{}) error {
	return DB.Model(&BuyerRequest{}).Where("request_id = ?", id).Updates(updatedFields).Error
}

// DeleteBuyerRequest 删除竞买人申请信息
func DeleteBuyerRequest(id string) error {
	return DB.Delete(&BuyerRequest{}, "request_id = ?", id).Error
}
func GetAllBuyerRequestWithConditions(conditions map[string]interface{}, page, pageSize int) ([]BuyerRequest, error) {
	var request []BuyerRequest

	// 计算 OFFSET
	offset := (page - 1) * pageSize

	// 使用 Limit 和 Offset 进行分页查询，并根据条件构建 WHERE 子句
	err := DB.Model(&request).Where(conditions).Limit(pageSize).Offset(offset).Find(&request).Error
	return request, err
}
