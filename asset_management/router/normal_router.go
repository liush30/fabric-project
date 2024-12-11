package router

import (
	"asset_management/handle"
	"asset_management/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterNormalRouter(r *gin.Engine) {
	normalRouter := r.Group("/normal")
	{
		normalRouter.POST("/buyer/apply", middleware.AuthMiddleware(), handle.BuyerApply)
		normalRouter.POST("/evaluator/apply", middleware.AuthMiddleware(), handle.AssessorApply)
		normalRouter.GET("/cancel/buyer/apply", handle.CancelBuyerRequest)
		normalRouter.GET("/cancel/evaluator/apply", handle.CancelAssessorRequest)
		normalRouter.POST("/query/buyer/apply", handle.QueryBuyerRequest)
		normalRouter.POST("/query/evaluator/apply", handle.QueryAssessorRequest)
	}
}
