package handler

import "github.com/gin-gonic/gin"

func SendErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{
		"message": message,
	})
}
