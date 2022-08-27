package util

import (
	"github.com/gin-gonic/gin"
)

func Respond(
	s int,
	p interface{},
	ctx *gin.Context,
) {
	ctx.JSON(
		s,
		p,
	)
}

func RespondError(
	s int,
	err error,
	ctx *gin.Context,
) {
	ctx.AbortWithError(
		s,
		err,
	)
}
