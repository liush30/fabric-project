package handle

import (
	"asset_management/db"
	"asset_management/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuctionTaskReq struct {
	AssetID     string `gorm:"column:asset_id"`     // 资产ID
	RuleID      string `gorm:"column:rule_id"`      // 拍卖规则ID
	TaskContent string `gorm:"column:task_content"` // 任务说明
	StartTime   string `gorm:"column:start_time"`   // 拍卖开始时间
	EndTime     string `gorm:"column:end_time"`     // 拍卖结束时间
	//Result      string `gorm:"column:result"`             // 拍卖结果
}

const (
	TaskStatePending = "待开始"
	TaskStateRunning = "进行中"
	TaskStateEnd     = "已结束"
)

// 创建拍卖任务
func CreateAuctionTask(c *gin.Context) {
	var req AuctionTaskReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
		return
	}
	task := db.AuctionTask{
		TaskID:      tool.GenerateUUIDWithoutDashes(),
		AssetID:     req.AssetID,
		RuleID:      req.RuleID,
		TaskContent: req.TaskContent,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		PublishTime: tool.GetNowTime(),
		TaskState:   TaskStatePending,
	}
	err = db.CreateAuctionTasks(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "创建成功"})
}

// UpdateAuctionTask 修改拍卖任务
func UpdateAuctionTask(c *gin.Context) {
	var task db.AuctionTask
	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
		return
	}

	err = db.UpdateAuctionTask(task.TaskID, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

// QueryAuctionTask 查询拍卖任务
func QueryAuctionTask(c *gin.Context) {
	page := c.Query("page")
	size := c.Query("size")
	conditions := make(map[string]interface{})
	err := c.ShouldBind(&conditions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
		return
	}
	// 将page和size转为int
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page转换失败:" + err.Error()})
		return
	}
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "size转换失败:" + err.Error()})
		return
	}
	info, err := db.GetAllAuctionTasksWithConditions(conditions, pageInt, sizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": info, "total": len(info)})
}

// UpdateAssetTaskStatus 修改任务状态
func UpdateAssetTaskStatus(c *gin.Context) {
	state := c.Query("task_state")
	taskId := c.Query("task_id")
	if state == "" || taskId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "state或asset_id不存在"})
		return
	}
	conditions := make(map[string]interface{})
	conditions["task_state"] = state
	err := db.UpdateAuctionTask(taskId, conditions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}
