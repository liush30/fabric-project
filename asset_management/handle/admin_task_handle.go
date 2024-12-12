package handle

import (
	"asset_management/db"
	"asset_management/fabric"
	"asset_management/tool"
	"encoding/json"
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

func QueryAuctionTaskByID(c *gin.Context) {
	taskId := c.Query("task_id")
	if taskId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "task_id不存在"})
		return
	}
	info, err := db.GetAuctionTaskByID(taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": info})
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

func StartAssetTask(c *gin.Context) {
	//获取task_id
	taskID := c.Query("task_id")
	condition := make(map[string]interface{})
	condition["task_state"] = TaskStateRunning
	err := db.UpdateAuctionTask(taskID, condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

type UpdateTaskStatusReq struct {
	TaskID string `json:"task_id"`
	//TaskState string `json:"task_state"`
	Result string `json:"result"`
	Notes  string `json:"notes"`
}

// EndAssetTaskStatus 修改任务状态
func EndAssetTaskStatus(c *gin.Context) {
	var req UpdateTaskStatusReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
		return
	}
	if req.TaskID == "" || req.Result == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "state或asset_id或result不存在"})
		return
	}
	conditions := make(map[string]interface{})
	conditions["task_state"] = TaskStateEnd

	err = db.UpdateAuctionTask(req.TaskID, conditions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败:" + err.Error()})
		return
	}
	//根据id查询
	task, err := db.GetAuctionTaskByID(req.TaskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	//计算task hash
	taskBytes, err := json.Marshal(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败:" + err.Error()})
		return
	}

	//任务结束
	err = fabric.CreateAuction(task.AssetID, req.TaskID, tool.CalculateSHA256Hash(taskBytes), req.Result, req.Notes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}
