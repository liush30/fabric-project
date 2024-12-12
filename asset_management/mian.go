package main

import (
	"asset_management/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type"}
	r.Use(cors.New(corsConfig))

	// 注册路由
	router.EvaluatorRouter(r)
	router.UserRouter(r)
	router.RegisterAdminRouter(r)
	router.RegisterNormalRouter(r)

	// 启动服务
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
