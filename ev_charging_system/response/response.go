package response

import (
	"ev_charging_system/errorcode"
	"ev_charging_system/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RespondOK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200})
}

func RespondErr(c *gin.Context, err string) {
	c.JSON(http.StatusOK, gin.H{"code": errorcode.ErrCodeCommon, "message": err})
}

func RespondDefaultErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": ""})
}

func RespondInvalidArgsErr(c *gin.Context, hint ...interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": errorcode.ErrCodeInvalidArgs, "message": "invalid args"})
	if len(hint) > 0 {
		log.Warnf("error argus:%v", hint[0])
	}

}

func RespondWithErrCode(c *gin.Context, code uint32, err string) {
	c.JSON(http.StatusOK, gin.H{"code": code, "message": err})
}

func RespondWithErrCode2(c *gin.Context, err error) {
	serr, ok := err.(*errorcode.ServiceError)
	if !ok {
		log.Warning("wrong error func call")
		RespondErr(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": serr.Code(), "message": serr.Error()})
}

func RespondWithData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"code": 1, "data": data})
}

func RespondServerError(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": errorcode.ErrCodeCommon, "message": "busy"})
}

func RespondWithRawJsonData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
