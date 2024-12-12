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

type AssetReq struct {
	AssetName    string `json:"asset_name"`
	AssetDate    string `json:"asset_date"`
	AssetContent string `json:"asset_content"`
	AssetOwner   string `json:"asset_owner"`
	Evaluator    string `json:"evaluator"`
}

// CreateAsset 上传资产信息
func CreateAsset(c *gin.Context) {
	var info AssetReq
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
	}
	//获取file
	file, err := c.FormFile("file")
	asset := db.AssetInfo{
		AssetID:         tool.GenerateUUIDWithoutDashes(),
		AssetName:       info.AssetName,
		AssetDate:       info.AssetDate,
		AssetContent:    info.AssetContent,
		AssetStatus:     db.AssetStatusPending,
		AssetOwner:      info.AssetOwner,
		CreateTime:      tool.GetNowTime(),
		Evaluator:       info.Evaluator,
		EvaluatorStatus: db.EvaluatorStatusPending,
	}
	if err == nil {
		path := "assets/" + file.Filename
		//保存文件
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败:" + err.Error()})
			return
		}
		asset.AssetImg = path
	}
	err = db.CreateAssetInfo(&asset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败:" + err.Error()})
		return
	}
	assetBytes, err := json.Marshal(asset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败:" + err.Error()})
		return
	}

	err = fabric.InitAsset(asset.AssetID, tool.CalculateSHA256Hash(assetBytes), asset.AssetOwner, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "资产上链失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "创建成功"})
}

// UpdateAsset 修改资产信息
func UpdateAsset(c *gin.Context) {
	var info db.AssetInfo
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
	}
	//获取file
	file, err := c.FormFile("file")
	if err == nil {
		path := "assets/" + file.Filename
		//保存文件
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败:" + err.Error()})
			return
		}
		info.AssetImg = path
	}
	err = db.UpdateAssetInfo(info.AssetID, info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败:" + err.Error()})
		return
	}
	assetBytes, err := json.Marshal(info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败:" + err.Error()})
		return
	}

	err = fabric.UpdateAsset(info.AssetID, tool.CalculateSHA256Hash(assetBytes), info.AssetOwner, "update asset info")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "资产上链失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

// QueryAssetByID 根据资产ID查询资产信息
func QueryAssetByID(c *gin.Context) {
	assetId := c.Query("asset_id")
	if assetId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "asset_id不存在"})
		return
	}
	info, err := db.GetAssetInfoByID(assetId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": info})
}

// QueryAssetInfo 查询资产信息
func QueryAssetInfo(c *gin.Context) {
	//获取page和size
	page := c.Query("page")
	size := c.Query("size")
	//判断page和size是否存在
	if page == "" || size == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "page或size不存在"})
		return
	}
	//将page和size转为int
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
	conditions := make(map[string]interface{})
	err = c.ShouldBind(&conditions)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
		return
	}
	info, err := db.GetAllAssetInfoWithConditions(conditions, pageInt, sizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": info, "total": len(info)})
}

// UpdateAssetStatus 修改资产状态
func UpdateAssetStatus(c *gin.Context) {
	state := c.Query("state")
	assetId := c.Query("asset_id")
	owner := c.Query("owner")
	if state == "" || assetId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "state或asset_id不存在"})
		return
	}
	conditions := make(map[string]interface{})
	conditions["asset_status"] = state
	err := db.UpdateAssetInfo(assetId, conditions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败:" + err.Error()})
		return
	}
	//计算hash
	id, err := db.GetAssetInfoByID(assetId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	assetBytes, err := json.Marshal(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败:" + err.Error()})
		return
	}
	err = fabric.UpdateAsset(assetId, tool.CalculateSHA256Hash(assetBytes), owner, "update status to "+state)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "资产上链失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

type AssetDetail struct {
	AssetInfo         db.AssetInfo              `json:"info"`
	AssetHistory      []fabric.Asset            `json:"history"`
	AuctionHistory    []fabric.Auction          `json:"auction_history"`
	AssessmentHistory []fabric.AssessmentResult `json:"assessments"`
	AuctionDetail     []AuctionDetail           `json:"auction_detail"`
}

type AuctionDetail struct {
	TaskID      string `json:"task_id"`      // 任务ID
	TaskState   string `json:"task_state"`   // 任务状态
	TaskContent string `json:"task_content"` // 任务说明
	StartTime   string `json:"start_time"`   // 拍卖开始时间
	EndTime     string `json:"end_time"`     // 拍卖结束时间
	PublishTime string `json:"publish_time"` // 发布时间
	RuleTitle   string `json:"rule_title"`   // 规则标题
	AssetHash   string `json:"hash"`
	Result      string `json:"result"`
	Notes       string `json:"notes"`
}

func QueryAssetDetail(c *gin.Context) {
	var detail AssetDetail
	//获取asset_id
	assetId := c.Query("asset_id")
	if assetId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "asset_id不存在"})
		return
	}
	info, err := db.GetAssetInfoByID(assetId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	detail.AssetInfo = *info
	history, err := fabric.QueryAssessmentHistory(assetId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询评估记录失败:" + err.Error()})
		return
	}
	detail.AssessmentHistory = history
	assetHistory, err := fabric.QueryAssetHistory(assetId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询资产历史记录失败:" + err.Error()})
		return
	}
	detail.AssetHistory = assetHistory
	auctionHistory, err := fabric.QueryAuctionHistory(assetId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询拍卖历史记录失败:" + err.Error()})
		return
	}
	for _, auction := range auctionHistory {
		//根据任务id查询任务
		task, err := db.GetAuctionTaskByID(auction.AuctionId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询拍卖任务失败:" + err.Error()})
			return
		}
		rule, err := db.GetAuctionRuleByID(task.RuleID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "查询拍卖规则失败:" + err.Error()})
			return
		}
		auctionDetail := AuctionDetail{
			TaskID:      task.TaskID,
			TaskState:   task.TaskState,
			TaskContent: task.TaskContent,
			StartTime:   task.StartTime,
			EndTime:     task.EndTime,
			PublishTime: task.PublishTime,
			RuleTitle:   rule.RuleTitle,
			AssetHash:   auction.AssetHash,
			Result:      auction.Result,
			Notes:       auction.Notes,
		}
		detail.AuctionDetail = append(detail.AuctionDetail, auctionDetail)
	}

	c.JSON(http.StatusOK, gin.H{"data": detail})

}
