package util

import (
	"github.com/gin-gonic/gin"
)

func Respond(
	c int,
	d interface{},
	ctx *gin.Context,
) {
	ctx.JSON(
		c,
		d,
	)
}

func RespondError(
	c int,
	err error,
	ctx *gin.Context,
) {
	ctx.AbortWithError(
		c,
		err,
	)
}
