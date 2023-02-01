package util

import (
	"github.com/gin-gonic/gin"
)

func SetContentType(
	ctx *gin.Context,
	contentType string,
) {
	ctx.Header("Content-Type", contentType)
}
