package handle

import (
	"asset_management/db"
	"asset_management/tool"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// QueryBuyerInfo 竞买人申请信息查询
func QueryBuyerInfo(c *gin.Context) {
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
	info, err := db.GetAllBuyerRequestWithConditions(conditions, pageInt, sizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": info, "total": len(info)})
}

// QueryAssessorInfo 评估者申请信息查询
func QueryAssessorInfo(c *gin.Context) {
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
	info, err := db.GetAllEvaluatorRequestWithConditions(conditions, pageInt, sizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": info, "total": len(info)})
}

type Audit struct {
	RequestID        string `json:"request_id"`
	Result           string `json:"result"`
	RejectionContent string `json:"rejection_content"`
}

// AuditBuyerRequest 审核竞买人申请
func AuditBuyerRequest(c *gin.Context) {
	var info Audit
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
	}
	condition := make(map[string]interface{})
	condition["result"] = info.Result
	condition["rejection_content"] = info.RejectionContent
	condition["process_time"] = tool.GetNowTime()
	err = db.UpdateBuyerRequest(info.RequestID, condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "审核失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "审核成功"})
}

// AuditAssessorRequest 审核评估者申请
func AuditAssessorRequest(c *gin.Context) {
	var info Audit
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
	}
	condition := make(map[string]interface{})
	condition["result"] = info.Result
	condition["rejection_content"] = info.RejectionContent
	condition["process_time"] = tool.GetNowTime()
	err = db.UpdateEvaluatorRequest(info.RequestID, condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "审核失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "审核成功"})
}
