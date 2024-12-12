package handle

import (
	"asset_management/db"
	"asset_management/fabric"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type EvaluatorResult struct {
	AssetId string `json:"asset_id"`
	Result  string `json:"result"`
	Note    string `json:"note"`
}

func CreateEvaluatorResult(c *gin.Context) {
	//获取userId
	userId := c.MustGet("userId").(string)
	var req EvaluatorResult
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
		return
	}
	condition := make(map[string]interface{})
	condition["evaluator_status"] = db.EvaluatorStatusEvaluated
	//修改资产状态
	err = db.UpdateAssetInfo(req.AssetId, condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败:" + err.Error()})
		return
	}
	err = fabric.UploadAssessmentResult(req.AssetId, userId, req.Result, req.Note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "上传失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "上传成功"})
}

// QueryEvaluatorResult 查询评估者所需评估资产
func QueryEvaluatorResult(c *gin.Context) {
	//获取userId
	userId := c.MustGet("userId").(string)
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
	info, err := db.GetAssetInfoByEvaluator(userId, pageInt, sizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": info, "total": len(info)})
}
