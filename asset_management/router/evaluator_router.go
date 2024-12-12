package router

import (
	"asset_management/handle"
	"asset_management/middleware"
	"github.com/gin-gonic/gin"
)

func EvaluatorRouter(r *gin.Engine) {
	//上传评估结果
	evaluatorRouter := r.Group("/evaluator")
	{
		evaluatorRouter.GET("/getAll", middleware.AuthMiddleware(), handle.QueryEvaluatorResult)
		evaluatorRouter.POST("/upload", middleware.AuthMiddleware(), handle.CreateEvaluatorResult)
	}
}
