package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maikonalexandre/govagas/handler"
	"github.com/maikonalexandre/govagas/middleware"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()
	middleware.InitializeHandler()

	v1 := router.Group("/api/v1")

	v1.Use(middleware.Auth)

	v1.GET("/health", handler.HealthHandler)

	v1.GET("/opening", handler.ShowOpeningHandler)
	v1.POST("/opening", handler.CreateOpeningHandler)
	v1.PUT("/opening", handler.UpdateOpeningHandler)
	v1.DELETE("/opening", handler.DeleteOpeningHandler)

	v1.GET("/openings", handler.ListOpeningHandler)
}
