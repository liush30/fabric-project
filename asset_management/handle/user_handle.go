package handle

import (
	"asset_management/db"
	"asset_management/tool"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginInfo struct {
	UserName string `json:"username"` //身份证号码
	Password string `json:"password"` //密码
}

func Login(c *gin.Context) {
	var info LoginInfo
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind:" + err.Error()})
	}
	user, err := db.Login(info.UserName, info.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to login:" + err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user does not exist"})
		return
	}
	token, err := tool.GenerateJWT(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "type": user.UserType})
}
