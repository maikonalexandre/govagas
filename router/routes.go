package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	v1.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	v1.POST("/opening", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	v1.PUT("/opening", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	v1.DELETE("/opening", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	v1.GET("/openings", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

}
