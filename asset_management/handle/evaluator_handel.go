package handle

import (
	"asset_management/fabric"
	"github.com/gin-gonic/gin"
	"net/http"
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
	err = fabric.UploadAssessmentResult(req.AssetId, userId, req.Result, req.Note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "上传失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "上传成功"})
}
