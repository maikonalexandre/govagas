package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maikonalexandre/govagas/schemas"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendErrorResponse(ctx, http.StatusBadRequest, errParamIsRequired("id", "query parameter").Error())
		return
	}

	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.ValidateUpdate(); err != nil {
		logger.Errf("validation error %v", err.Error())
		SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, &id).Error; err != nil {
		SendErrorResponse(ctx, http.StatusNotFound, fmt.Sprintf("opening whith id: %s not found", id))
		return
	}

	if err := db.Model(&opening).Updates(&request).Error; err != nil {
		SendErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("error on update opening with id: %s", id))
	}

	SendSuccesResponse(ctx, "delete-opening", &opening)
}
