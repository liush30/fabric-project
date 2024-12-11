package db

import "fmt"

type EvaluatorRequest struct {
	RequestID        string `gorm:"primaryKey;column:request_id"` // 请求ID
	UserID           string `gorm:"column:user_id"`               // 用户ID
	CertNumber       string `gorm:"column:cert_number"`           // 资历证书编号
	QuaContent       string `gorm:"column:qua_content"`           // 资历说明
	EvaluatorName    string `gorm:"column:evaluator_name"`        // 评估者姓名
	EvaluatorPhone   string `gorm:"column:evaluator_phone"`       // 评估者联系方式
	RejectionContent string `gorm:"column:rejection_content"`     // 被拒原因
	Result           string `gorm:"column:result"`                // 认证结果
	RequestTime      string `gorm:"column:request_time"`          // 申请时间
	ProcessTime      string `gorm:"column:process_time"`          // 处理时间
}

func (EvaluatorRequest) TableName() string {
	return "evaluator_request"
}

func CreateEvaluatorRequest(EvaluatorRequest *EvaluatorRequest) error {
	result := DB.Create(EvaluatorRequest)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("insert failed,affected rows:%d", result.RowsAffected)
	}
	return nil
}

// GetEvaluatorRequestByID 查询评估者申请信息
func GetEvaluatorRequestByID(id string) (*EvaluatorRequest, error) {
	var EvaluatorRequest EvaluatorRequest
	err := DB.Model(&EvaluatorRequest).First(&EvaluatorRequest, "request_id = ?", id).Error
	return &EvaluatorRequest, err
}

// UpdateEvaluatorRequest 更新评估者申请信息
func UpdateEvaluatorRequest(id string, updatedFields interface{}) error {
	return DB.Model(&EvaluatorRequest{}).Where("request_id = ?", id).Updates(updatedFields).Error
}

// DeleteEvaluatorRequest 删除评估者申请信息
func DeleteEvaluatorRequest(id string) error {
	return DB.Delete(&EvaluatorRequest{}, "request_id = ?", id).Error
}
func GetAllEvaluatorRequestWithConditions(conditions map[string]interface{}, page, pageSize int) ([]EvaluatorRequest, error) {
	var asset []EvaluatorRequest

	// 计算 OFFSET
	offset := (page - 1) * pageSize

	// 使用 Limit 和 Offset 进行分页查询，并根据条件构建 WHERE 子句
	err := DB.Model(&asset).Where(conditions).Limit(pageSize).Offset(offset).Find(&asset).Error
	return asset, err
}
