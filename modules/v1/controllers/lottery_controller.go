package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LotteryController struct {
	BaseController
}

func (this *LotteryController) GetLotteryList(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"msg": "hahaha",
	})
}
