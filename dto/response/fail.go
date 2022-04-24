package responseDto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

func FailWithParams(ctx *gin.Context) {
	ctx.JSON(http.StatusUnprocessableEntity, errorResponse{
		ErrorCode: "0001",
		Message:   "please checking params are all correct.",
	})
}

func FailWithError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, errorResponse{
		ErrorCode: "0002",
		Message:   err.Error(),
	})
}
