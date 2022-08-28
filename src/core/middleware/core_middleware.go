package middleware

import "github.com/gin-gonic/gin"

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()
	if len(ctx.Errors) > 0 {
		ctx.JSON(
			-1,
			gin.H{"error": ctx.Errors[0].Error()},
		)
	}
}
