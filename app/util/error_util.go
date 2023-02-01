package util

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func ErrorAbort(
	code int,
	message string,
	ctx *gin.Context,
) {
	ctx.AbortWithError(
		code,
		errors.New(
			message,
		),
	)
}
