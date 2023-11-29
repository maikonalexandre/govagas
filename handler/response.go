package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, gin.H{
		"message": message,
	})
}

func SendSuccesResponse(ctx *gin.Context, op string, data ...interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%s successfull", op),
		"data":    data,
	})
}
