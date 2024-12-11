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

const (
	UserTypeBuyer     = "竟买者"
	UserTypeEvaluator = "评估者"
	UserTypeAdmin     = "管理员"
	UserTypeNormal    = "普通用户"
)

func Login(c *gin.Context) {
	var info LoginInfo
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
	}
	user, err := db.Login(info.UserName, info.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "登录失败:" + err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户不存在"})
		return
	}
	token, err := tool.GenerateJWT(user.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成token失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "type": user.UserType, "message": "登录成功！"})
}

func Register(c *gin.Context) {
	var info LoginInfo
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败:" + err.Error()})
	}
	user := db.UserInfo{
		UserID:       tool.GenerateUUIDWithoutDashes(),
		UserPassword: info.Password,
		UserType:     UserTypeNormal,
		UserName:     info.UserName,
		RegisterTime: tool.GetNowTime(),
	}
	err = db.CreateUserInfo(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "注册成功！"})
}
