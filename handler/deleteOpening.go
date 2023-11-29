package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maikonalexandre/govagas/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		SendErrorResponse(ctx, http.StatusBadRequest, errParamIsRequired("id", "query parameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, &id); err != nil {
		SendErrorResponse(ctx, http.StatusNotFound, fmt.Sprintf("opening whith id: %s not found", id))
		return
	}

	if err := db.Delete(&opening, &id); err != nil {
		SendErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("error on delete opening whith id: %s", id))
		return
	}

	SendSuccesResponse(ctx, "delete-opening", &opening)
}
