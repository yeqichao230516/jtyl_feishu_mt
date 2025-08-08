package controller

import "github.com/gin-gonic/gin"

func ZntzOut(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "success",
	})
}
func ZntzIn(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "success",
	})
}
