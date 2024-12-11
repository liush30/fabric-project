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
	AssetOwner   string `gorm:"column:asset_owner"`
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
		AssetID:      tool.GenerateUUIDWithoutDashes(),
		AssetName:    info.AssetName,
		AssetDate:    info.AssetDate,
		AssetContent: info.AssetContent,
		AssetStatus:  db.AssetStatusPending,
		AssetOwner:   info.AssetOwner,
		CreateTime:   tool.GetNowTime(),
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
