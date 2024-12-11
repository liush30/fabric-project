package db

import (
	"fmt"
)

type AuctionTask struct {
	TaskID      string `gorm:"primaryKey;column:task_id"` // 任务ID
	AssetID     string `gorm:"column:asset_id"`           // 资产ID
	RuleID      string `gorm:"column:rule_id"`            // 拍卖规则ID
	TaskState   string `gorm:"column:task_state"`         // 任务状态
	TaskContent string `gorm:"column:task_content"`       // 任务说明
	StartTime   string `gorm:"column:start_time"`         // 拍卖开始时间
	EndTime     string `gorm:"column:end_time"`           // 拍卖结束时间
	PublishTime string `gorm:"column:publish_time"`       // 发布时间
	//Result      string `gorm:"column:result"`             // 拍卖结果
}

func (AuctionTask) TableName() string {
	return "auction_task"
}
func CreateAuctionTasks(AuctionTasks *AuctionTask) error {
	result := DB.Create(AuctionTasks)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("insert failed,affected rows:%d", result.RowsAffected)
	}
	return nil
}

// GetAuctionTaskByID 查询拍卖任务
func GetAuctionTaskByID(id string) (*AuctionTask, error) {
	var AuctionTask AuctionTask
	err := DB.Model(&AuctionTask).First(&AuctionTask, "task_id = ?", id).Error
	return &AuctionTask, err
}

// UpdateAuctionTask 更新拍卖任务
func UpdateAuctionTask(id string, updatedFields interface{}) error {
	return DB.Model(&AuctionTask{}).Where("task_id = ?", id).Updates(updatedFields).Error
}

// DeleteAuctionTask 删除拍卖任务
func DeleteAuctionTask(id string) error {
	return DB.Delete(&AuctionTask{}, "task_id = ?", id).Error
}
func GetAllAuctionTasksWithConditions(conditions map[string]interface{}, page, pageSize int) ([]AuctionTask, error) {
	var task []AuctionTask

	// 计算 OFFSET
	offset := (page - 1) * pageSize

	// 使用 Limit 和 Offset 进行分页查询，并根据条件构建 WHERE 子句
	err := DB.Model(&task).Where(conditions).Limit(pageSize).Offset(offset).Find(&task).Error
	return task, err
}
