package content

import (
	"github.com/gin-gonic/gin"
)

type ContentHandler struct {
	service *ContentService
}

func (h *ContentHandler) InitDependencies() {
	repo := &ContentRepository{}
	h.service = &ContentService{
		repo: repo,
	}
}

func (h *ContentHandler) Create(ctx *gin.Context) {
	h.service.Create(
		ctx,
	)
}

func (h *ContentHandler) FindAll(ctx *gin.Context) {
	h.service.FindAll(
		ctx,
	)
}

func (h *ContentHandler) FindOne(ctx *gin.Context) {
	h.service.FindOne(
		ctx,
	)
}

func (h *ContentHandler) Update(ctx *gin.Context) {
	h.service.Update(
		ctx,
	)
}

func (h *ContentHandler) Delete(ctx *gin.Context) {
	h.service.Delete(
		ctx,
	)
}
