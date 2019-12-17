package main

import (
	"gin_lottery/app"
	router2 "gin_lottery/modules/v1/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){

	app.InitApp()

	router := gin.Default()

	//注册路由
	router2.V1InitRouter(router)

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "hello",
		})
	})

	router.Run(":9999")

}
