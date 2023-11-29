package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maikonalexandre/govagas/config"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *config.Logger
)

func InitializeHandler() {
	logger = config.GetLogger("handler: ")
}

func HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "ok"})
}
