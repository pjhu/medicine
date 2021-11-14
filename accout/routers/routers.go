package routers

import (
	"github.com/gin-gonic/gin"
)

// SetupRouterEngine 初始化
func SetupRouterEngine() *gin.Engine {

	r := gin.Default()

	r.POST("/api/v1/accounts/decrease", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"userId": "user",
			"orderAmount": 1,
		})
	})
	
	r.GET("/api/v1/accounts/:id", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"orderAmount": 100,
		})
	})
	return r
}