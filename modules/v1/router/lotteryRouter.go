package router

import (
	"gin_lottery/modules/v1/controllers"
	"github.com/gin-gonic/gin"
)

func lottery(router *gin.RouterGroup){

	lotteryController := &controllers.LotteryController{}

	groupRouter := router.Group("/lottery")

	//获取首页抽奖列表
	groupRouter.GET("/getLotteryList", lotteryController.GetLotteryList)
}
