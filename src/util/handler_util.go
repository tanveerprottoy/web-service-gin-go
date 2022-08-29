package util

import (
	"github.com/gin-gonic/gin"
)

func HandleReq(
	c int,
	p interface{},
	err error,
	ctx *gin.Context,
) {
	if err != nil {
		RespondError(
			c,
			err,
			ctx,
		)
		return
	}
	Respond(
		c,
		p,
		ctx,
	)
}
