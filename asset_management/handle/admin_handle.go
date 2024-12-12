package handle

import (
	"asset_management/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func QueryUserInfo(c *gin.Context) {
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
	info, err := db.GetAllUserInfoWithConditions(conditions, pageInt, sizeInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": info, "total": len(info)})
}

// UpdateUserInfo 修改用户信息
func UpdateUserInfo(c *gin.Context) {
	var info db.UserInfo
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
	}
	err = db.UpdateUserInfo(info.UserID, info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "更新成功"})
}

// 获取指定用户的信息
func QueryUserInfoByID(c *gin.Context) {
	userId := c.Query("user_id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id不存在"})
		return
	}
	info, err := db.GetUserInfoByID(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": info})
}

// DeleteUserInfo 删除用户信息
func DeleteUserInfo(c *gin.Context) {
	userId := c.Query("user_id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id不存在"})
		return
	}
	err := db.DeleteUserInfo(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

// QueryAllAuctionType 查询所有拍卖类型
func QueryAllAuctionType(c *gin.Context) {
	ruleType, err := db.QueryAllAuctionRuleType()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ruleType})
}
