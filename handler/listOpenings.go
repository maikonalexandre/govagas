package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maikonalexandre/govagas/schemas"
)

func ListOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		SendErrorResponse(ctx, http.StatusInternalServerError, "error on lists opnings")
		return
	}

	SendSuccesResponse(ctx, "list-openings", openings)
}
