package middleware

import (
	"github.com/gin-gonic/gin"
)

func JSONMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Next()
}

func ErrorMiddleware(ctx *gin.Context) {
	ctx.Next()
	if len(ctx.Errors) > 0 {
		ctx.JSON(
			-1,
			map[string]interface{}{"error": ctx.Errors[0].Error()},
		)
	}
}
