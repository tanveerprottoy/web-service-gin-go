package util

import (
	"github.com/gin-gonic/gin"
)

func Respond(
	s int,
	p interface{},
	c *gin.Context,
) {
	c.JSON(
		s,
		p,
	)
}

func RespondError(
	s int,
	err error,
	c *gin.Context,
) {
	c.AbortWithError(
		s,
		err,
	)
}
