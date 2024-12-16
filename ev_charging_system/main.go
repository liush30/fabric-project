package main

import (
	"ev_charging_system/config"
	"ev_charging_system/log"

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

	// 启动服务
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func handleRecovery(c *gin.Context, err interface{}) {
	c.JSON(1010, "internal server error")
}
