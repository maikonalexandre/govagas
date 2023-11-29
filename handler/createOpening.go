package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("validation error %v", err.Error())
		SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.Create(&request).Error; err != nil {
		logger.Errf("error creating opening %v", err.Error())
		return
	}
}
