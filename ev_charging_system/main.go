package main

import (
	"ev_charging_system/config"
	"ev_charging_system/controller"
	"ev_charging_system/log"
	"ev_charging_system/middleware"
	"runtime"

	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	//加载配置文件
	sysType := runtime.GOOS
	switch sysType {
	case "windows", "darwin":
		config.LoadConfig("config/config.yaml")
	case "linux":
		if len(os.Args) < 1 {
			panic("miss config file")
		}
		config.LoadConfig(os.Args[1])
	default:
		panic(fmt.Sprintf("error system:%s", sysType))
	}
	log.InitLogger("web") //初始化日志器
	r := gin.Default()
	if config.ChargeConfig.NodeEnv == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type"}
	r.Use(cors.New(corsConfig))
	r.Use(gin.RecoveryWithWriter(log.LoggerWriter(), handleRecovery), gin.LoggerWithWriter(log.LoggerWriter()))
	//tool.GenDao()
	user := r.Group("/user")
	user.POST("/login", controller.RepairmanController.Login)
	user.Use(middleware.AuthMiddleware())
	user.GET("/info", controller.RepairmanController.Info)
	user.POST("/page", controller.RepairmanController.ListAndPage)
	user.GET("/info/:userId", controller.RepairmanController.GetUserById)
	user.POST("/add", controller.RepairmanController.AddUser)
	user.POST("/update", controller.RepairmanController.UpdateUser)

	// 启动服务
	if err := r.Run(fmt.Sprintf("0.0.0.0:%s", config.ChargeConfig.WebInfo.Port)); err != nil {
		panic(err)
	}
}

func handleRecovery(c *gin.Context, err interface{}) {
	c.JSON(500, "internal server error")
}
