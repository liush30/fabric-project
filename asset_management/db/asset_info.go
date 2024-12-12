package db

import "fmt"

type AssetInfo struct {
	AssetID         string `gorm:"primaryKey;column:asset_id"` // 资产ID
	AssetStatus     string `gorm:"column:asset_status"`        // 资产当前状态
	AssetName       string `gorm:"column:asset_name"`          // 资产名称
	AssetDate       string `gorm:"column:asset_date"`          // 资产产出时间
	AssetContent    string `gorm:"column:asset_content"`       // 资产介绍
	AssetImg        string `gorm:"column:asset_img"`           // 资产图片
	AssetOwner      string `gorm:"column:asset_owner"`         // 资产所有者ID(初始)
	CreateTime      string `gorm:"column:create_time"`         // 记录创建时间
	Evaluator       string `gorm:"column:evaluator"`           // 评估者
	EvaluatorStatus string `gorm:"column:evaluator_status"`    // 评估状态
}

const (
	AssetStatusPending = "待上架"
	//上架中、拍卖中、已流拍、已售出、已下架
	AssetStatusOnShelf       = "上架中"
	AssetStatusAuction       = "拍卖中"
	AssetStatusFail          = "已流拍"
	AssetStatusSoldOut       = "已售出"
	AssetStatusOffShelf      = "已下架"
	EvaluatorStatusPending   = "待评估"
	EvaluatorStatusEvaluated = "已评估"
)

func (a *AssetInfo) TableName() string {
	return "asset_info"
}
func CreateAssetInfo(AssetInfo *AssetInfo) error {
	result := DB.Create(AssetInfo)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("insert failed,affected rows:%d", result.RowsAffected)
	}
	return nil
}

// GetAssetInfoByID 查询资产信息
func GetAssetInfoByID(id string) (*AssetInfo, error) {
	var AssetInfo AssetInfo
	err := DB.Model(&AssetInfo).First(&AssetInfo, "asset_id = ?", id).Error
	return &AssetInfo, err
}

// UpdateAssetInfo 更新资产信息
func UpdateAssetInfo(id string, updatedFields interface{}) error {
	return DB.Model(&AssetInfo{}).Where("asset_id = ?", id).Updates(updatedFields).Error
}

// DeleteAssetInfo 删除资产信息
func DeleteAssetInfo(id string) error {
	return DB.Delete(&AssetInfo{}, "asset_id = ?", id).Error
}
func GetAllAssetInfoWithConditions(conditions map[string]interface{}, page, pageSize int) ([]AssetInfo, error) {
	var asset []AssetInfo

	// 计算 OFFSET
	offset := (page - 1) * pageSize

	// 使用 Limit 和 Offset 进行分页查询，并根据条件构建 WHERE 子句
	err := DB.Model(&asset).Where(conditions).Limit(pageSize).Offset(offset).Find(&asset).Error
	return asset, err
}

// 查询指定评估者的资产
func GetAssetInfoByEvaluator(evaluator string, page, pageSize int) ([]AssetInfo, error) {
	var assets []AssetInfo
	err := DB.Model(&AssetInfo{}).Where("evaluator = ?", evaluator).Limit(pageSize).Offset((page - 1) * pageSize).Find(&assets).Error
	return assets, err
}
