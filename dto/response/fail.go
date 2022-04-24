package responseDto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Code    string `json:"Code"`
	Message string `json:"message"`
}

func FailWithParams(ctx *gin.Context) {
	ctx.JSON(http.StatusUnprocessableEntity, errorResponse{
		Code:    "01",
		Message: "please checking params are all correct.",
	})
}

func FailWithError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, errorResponse{
		Code:    "02",
		Message: err.Error(),
	})
}
