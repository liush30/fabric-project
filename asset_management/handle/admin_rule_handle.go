package handle

import (
	"asset_management/db"
	"asset_management/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuctionRuleReq struct {
	RuleTitle    string `json:"rule_title"`    // 规则标题
	RuleContent  string `json:"rule_content"`  // 规则说明
	BidIncrement string `json:"bid_increment"` // 最小加价增量
	MaxBidCount  string `json:"max_bid_count"` // 最大出价次数
	StartPrice   string `json:"start_price"`   // 起拍价
	TimeLimit    string `json:"time_limit"`    // 拍卖时限
	AutoExtend   string `json:"auto_extend"`   // 是否启用自动延时
	TypeID       string `json:"type_id"`       // 拍卖类型ID
	EnableAsset  string `json:"enable_asset"`  // 是否启用资产限制
	AssetLimit   int    `json:"asset_limit"`   // 资产限制值
	EnableCredit string `json:"enable_credit"` // 是否启用信用评分限制
	ScoreLimit   int    `json:"score_limit"`   // 信用评分限制值
}

func CreateAuctionRule(c *gin.Context) {
	var req AuctionRuleReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
		return
	}
	rule := db.AuctionRule{
		RuleID:       tool.GenerateUUIDWithoutDashes(),
		RuleTitle:    req.RuleTitle,
		RuleContent:  req.RuleContent,
		BidIncrement: req.BidIncrement,
		MaxBidCount:  req.MaxBidCount,
		StartPrice:   req.StartPrice,
		TimeLimit:    req.TimeLimit,
		AutoExtend:   req.AutoExtend,
		TypeID:       req.TypeID,
		EnableAsset:  req.EnableAsset,
		AssetLimit:   req.AssetLimit,
		EnableCredit: req.EnableCredit,
		ScoreLimit:   req.ScoreLimit,
	}
	err = db.CreateAuctionRule(&rule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "创建成功"})
}

// UpdateAuctionRule  修改拍卖规则
func UpdateAuctionRule(c *gin.Context) {
	var rule db.AuctionRule
	err := c.ShouldBind(&rule)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
		return
	}

	err = db.UpdateAuctionRule(rule.RuleID, &rule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

// DeleteAuctionRule 删除拍卖规则
func DeleteAuctionRule(c *gin.Context) {
	ruleId := c.Query("rule_id")
	if ruleId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rule_id不存在"})
		return
	}
	err := db.DeleteAuctionRule(ruleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// QueryAuctionRule 查询拍卖规则
func QueryAuctionRule(c *gin.Context) {
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
	info, err := db.GetAllAuctionRuleWithConditions(conditions, pageInt, sizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": info, "total": len(info)})
}

func QueryAuctionRuleByID(c *gin.Context) {
	ruleId := c.Query("rule_id")
	if ruleId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "rule_id不存在"})
		return
	}
	info, err := db.GetAuctionRuleByID(ruleId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": info})
}
