package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maikonalexandre/govagas/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("validation error %v", err.Error())
		SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errf("error creating opening %v", err.Error())
		SendErrorResponse(ctx, http.StatusInternalServerError, "failed on creating database")
		return
	}

	SendSuccesResponse(ctx, "create-opening", &opening)
}
