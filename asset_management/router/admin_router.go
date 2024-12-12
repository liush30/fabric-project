package router

import (
	"asset_management/handle"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRouter(r *gin.Engine) {
	adminRouter := r.Group("/admin")
	{
		assetRouter := adminRouter.Group("/asset")
		{
			assetRouter.POST("/create", handle.CreateAsset)
			assetRouter.POST("/getAll", handle.QueryAssetInfo)
			assetRouter.POST("/update", handle.UpdateAsset)
			assetRouter.POST("/update/status", handle.UpdateAssetStatus)
			assetRouter.POST("/getDetail", handle.QueryAssetDetail)
			assetRouter.GET("/getOne", handle.QueryAssetByID)

		}
		taskRouter := adminRouter.Group("/task")
		{
			taskRouter.POST("/create", handle.CreateAuctionTask)
			taskRouter.POST("/getAll", handle.QueryAuctionTask)
			taskRouter.POST("/update", handle.UpdateAuctionTask)
			taskRouter.GET("/getOne", handle.QueryAuctionTaskByID)
			taskRouter.GET("/start", handle.StartAssetTask)
			taskRouter.POST("/end", handle.EndAssetTaskStatus)

		}
		ruleRouter := adminRouter.Group("/rule")
		{
			ruleRouter.POST("/create", handle.CreateAuctionRule)
			ruleRouter.POST("/getAll", handle.QueryAuctionRule)
			ruleRouter.POST("/update", handle.UpdateAuctionRule)
			ruleRouter.GET("/delete", handle.DeleteAuctionRule)
			ruleRouter.GET("/getOne", handle.QueryAuctionRuleByID)
			ruleRouter.GET("/getType", handle.QueryAllAuctionType)

		}
		requestRouter := adminRouter.Group("/request")
		{
			requestRouter.POST("/buyer/getAll", handle.QueryBuyerInfo)
			requestRouter.POST("/evaluator/getAll", handle.QueryAssessorInfo)
			requestRouter.GET("/buyer/cancel", handle.CancelBuyerRequest)
			requestRouter.GET("/evaluator/cancel", handle.CancelAssessorRequest)
			requestRouter.POST("/buyer/audit", handle.AuditBuyerRequest)
			requestRouter.POST("/evaluator/audit", handle.AuditAssessorRequest)

		}
		userRouter := adminRouter.Group("/user")
		{
			userRouter.POST("/getAll", handle.QueryUserInfo)
			userRouter.POST("/update", handle.UpdateUserInfo)
			userRouter.GET("/delete", handle.DeleteUserInfo)
			userRouter.GET("/getOne", handle.QueryUserInfoByID)
		}
	}

}
