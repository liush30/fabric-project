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
	user := r.Group("user")
	user.POST("/login", controller.RepairmanController.Login)
	user.Use(middleware.AuthMiddleware())
	user.GET("/info", controller.RepairmanController.Info)
	user.POST("/page", controller.RepairmanController.ListAndPage)
	user.GET("/info/:userId", controller.RepairmanController.GetUserById)
	user.POST("/add", controller.RepairmanController.AddUser)
	user.POST("/update", controller.RepairmanController.UpdateUser)
	user.GET("/delete/:userId", controller.RepairmanController.DeleteUser)

	station := r.Group("station")
	station.Use(middleware.AuthMiddleware())
	station.POST("/page", controller.StationController.StationByPage)
	station.POST("/add", controller.StationController.AddStation)
	station.POST("/update", controller.StationController.UpdateStation)
	station.GET("/info/:stationId", controller.StationController.GetStationById)
	station.GET("/info", controller.StationController.GetMeStationInfo)
	station.GET("/delete/:stationId", controller.StationController.DeleteStation)

	pile := r.Group("pile")
	pile.Use(middleware.AuthMiddleware())
	pile.POST("/page", controller.PileController.PileByPage)
	pile.POST("/add", controller.PileController.AddPile)
	pile.POST("/update", controller.PileController.UpdatePile)
	pile.GET("/info/:pileId", controller.PileController.GetPileById)
	pile.GET("/delete/:pileId", controller.PileController.DeletePile)
	pile.GET("/history/:pileId", controller.PileController.GetPileHistory)
	pile.POST("/me/page", controller.PileController.GetMePilePage)
	pile.POST("/me/add", controller.PileController.AddMePile)

	repairRequest := r.Group("repair")
	repairRequest.Use(middleware.AuthMiddleware())
	repairRequest.POST("/page", controller.RepairRequestController.RepairRequestByPage)
	repairRequest.POST("/add", controller.RepairRequestController.AddRepairRequest)
	repairRequest.POST("/update", controller.RepairRequestController.UpdateRepairRequest)
	repairRequest.GET("/cancel/:repairRequestId", controller.RepairRequestController.CancelRepairRequest)
	repairRequest.GET("/info/:repairRequestId", controller.RepairRequestController.GetRepairRequestById)
	repairRequest.POST("/me/add", controller.RepairRequestController.AddMeRepairRequest)

	fee := r.Group("fee")
	fee.Use(middleware.AuthMiddleware())
	fee.POST("/page", controller.FeeController.FeeByPage)
	fee.POST("/add", controller.FeeController.AddFee)
	fee.POST("/update", controller.FeeController.UpdateFee)
	fee.GET("/info/:FeeId", controller.FeeController.GetFeeById)
	fee.GET("/delete/:FeeId", controller.FeeController.DeleteFee)

	gun := r.Group("gun")
	gun.POST("/page", controller.GunController.GunByPage)
	gun.Use(middleware.AuthMiddleware())
	gun.POST("/add", controller.GunController.AddGun)
	gun.POST("/update", controller.GunController.UpdateGun)
	gun.GET("/info/:gunId", controller.GunController.GetGunById)
	gun.GET("/delete/:gunId", controller.GunController.DeleteGun)
	gun.GET("/list/pile/:pileId", controller.GunController.GetGunListByPileId)
	gun.GET("/historyByPile/:pileId", controller.GunController.GetGunHistoryByPileId)
	gun.GET("/history/:gunId", controller.GunController.GetGunHistory)

	// 启动服务
	if err := r.Run(fmt.Sprintf("0.0.0.0:%s", config.ChargeConfig.WebInfo.Port)); err != nil {
		panic(err)
	}
}

func handleRecovery(c *gin.Context, err interface{}) {
	c.JSON(500, "internal server error")
}
