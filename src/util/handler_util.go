package util

import (
	"github.com/gin-gonic/gin"
)

func HandleReq(
	code int,
	payload interface{},
	err error,
	c *gin.Context,
) {
	if err != nil {
		RespondError(
			code,
			err,
			c,
		)
		return
	}
	Respond(
		code,
		payload,
		c,
	)
}
