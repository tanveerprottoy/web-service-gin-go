package util

import (
	"github.com/gin-gonic/gin"
)

func Respond(code int, data any, ctx *gin.Context) {
	ctx.JSON(code, data)
}

func RespondError(code int, err error, ctx *gin.Context) {
	ctx.AbortWithError(code, err)
}
