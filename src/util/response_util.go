package util

import (
	"github.com/gin-gonic/gin"
)

func Respond(
	c int,
	p interface{},
	ctx *gin.Context,
) {
	ctx.JSON(
		c,
		p,
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
