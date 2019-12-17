package router

import "github.com/gin-gonic/gin"

func V1InitRouter(router *gin.Engine) *gin.Engine{

	v1 := router.Group("/v1")

	lottery(v1)

	return router
}
