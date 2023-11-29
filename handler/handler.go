package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func ShowOpeningHandler(ctx *gin.Context) {

}

func CreateOpeningHandler(ctx *gin.Context) {

}

func DeleteOpeningHandler(ctx *gin.Context) {

}

func UpdateOpeningHandler(ctx *gin.Context) {

}

func ListOpeningHandler(ctx *gin.Context) {

}
