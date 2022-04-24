package responseDto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
	return
}

type successMessage struct {
	Message string `json:"message"`
}

func SuccessWithMessage(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, successMessage{Message: msg})
	return
}

type successData struct {
	Result interface{} `json:"result"`
}

func SuccessWithData(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, successData{
		Result: data,
	})
}

func SuccessCreated(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusCreated, successData{
		Result: data,
	})
}
