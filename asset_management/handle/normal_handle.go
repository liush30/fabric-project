package handle

import (
	"asset_management/db"
	"asset_management/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BuyerReq struct {
	BuyerName    string `gorm:"column:buyer_name"`    // 竞买人姓名
	BuyerPhone   string `gorm:"column:buyer_phone"`   // 竞买人联系方式
	BuyerNumber  string `gorm:"column:buyer_number"`  // 竞买人身份证号码
	BuyerAddress string `gorm:"column:buyer_address"` // 竞买人地址
	TotalAssets  int    `gorm:"column:total_assets"`  // 总资产
	CreditScore  int    `gorm:"column:credit_score"`  // 信用评分
	IncomeSource string `gorm:"column:income_source"` // 收入来源
}

const (
	RequestStateCancel = "已取消"
)

// BuyerApply 竞买者申请
func BuyerApply(c *gin.Context) {
	var req BuyerReq
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
		return
	}
	userId := c.MustGet("userId").(string)
	buyerReq := db.BuyerRequest{
		RequestID:    tool.GenerateUUIDWithoutDashes(),
		UserID:       userId,
		BuyerName:    req.BuyerName,
		BuyerPhone:   req.BuyerPhone,
		BuyerNumber:  req.BuyerNumber,
		BuyerAddress: req.BuyerAddress,
		TotalAssets:  req.TotalAssets,
		CreditScore:  req.CreditScore,
		IncomeSource: req.IncomeSource,
		RequestTime:  tool.GetNowTime(),
	}
	err = db.CreateBuyerRequest(&buyerReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "创建成功"})
}

// QueryBuyerRequest 查询请求者发起的请求记录
func QueryBuyerRequest(c *gin.Context) {
	userId := c.MustGet("userId").(string)

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
	condition := make(map[string]interface{})
	condition["user_id"] = userId
	info, err := db.GetAllBuyerRequestWithConditions(conditions, pageInt, sizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": info, "total": len(info)})
}

type EvaluatorRequest struct {
	CertNumber     string `gorm:"column:cert_number"`     // 资历证书编号
	QuaContent     string `gorm:"column:qua_content"`     // 资历说明
	EvaluatorName  string `gorm:"column:evaluator_name"`  // 评估者姓名
	EvaluatorPhone string `gorm:"column:evaluator_phone"` // 评估者联系方式
}

// AssessorApply 评估者申请
func AssessorApply(c *gin.Context) {
	var req EvaluatorRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
		return
	}
	//获取userId
	userId := c.MustGet("userId").(string)
	assessorReq := db.EvaluatorRequest{
		RequestID:      tool.GenerateUUIDWithoutDashes(),
		UserID:         userId,
		CertNumber:     req.CertNumber,
		QuaContent:     req.QuaContent,
		EvaluatorName:  req.EvaluatorName,
		EvaluatorPhone: req.EvaluatorPhone,
		RequestTime:    tool.GetNowTime(),
	}
	err = db.CreateEvaluatorRequest(&assessorReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "创建成功"})
}

// QueryAssessorRequest 查询评估者发起的请求记录
func QueryAssessorRequest(c *gin.Context) {
	userId := c.MustGet("userId").(string)
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
	condition := make(map[string]interface{})
	condition["user_id"] = userId
	info, err := db.GetAllEvaluatorRequestWithConditions(conditions, pageInt, sizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": info, "total": len(info)})
}

// 取消竞买人认证申请
func CancelBuyerRequest(c *gin.Context) {
	//获取id
	request_id := c.Query("request_id")
	//获取id
	condition := make(map[string]interface{})
	condition["result"] = RequestStateCancel
	err := db.UpdateBuyerRequest(request_id, condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "取消失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "取消成功"})
}

// CancelAssessorRequest 取消评估人认证申请
func CancelAssessorRequest(c *gin.Context) {
	//获取id
	request_id := c.Query("request_id")
	//获取id
	condition := make(map[string]interface{})
	condition["result"] = RequestStateCancel
	err := db.UpdateEvaluatorRequest(request_id, condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "取消失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "取消成功"})
}
