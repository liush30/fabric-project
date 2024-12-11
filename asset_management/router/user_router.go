package router

import (
	"asset_management/handle"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	UserRouter := r.Group("/user")
	{
		UserRouter.POST("/login", handle.Login)
		UserRouter.POST("/register", handle.Register)
	}
}
