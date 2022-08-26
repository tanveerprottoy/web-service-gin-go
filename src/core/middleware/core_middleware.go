package middleware

import (
	"net/http"
	"txp/web-service-gin/src/util"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context) {
	ctx.Next()
	if len(ctx.Errors) > 0 {
		util.RespondError(
http.StatusInternalServerError,
		)
		ctx.JSON(
			,
			util.InternalServerError,
		)
	}
	ctx.JSON(
		http.StatusInternalServerError,
		util.InternalServerError,
	)
}
