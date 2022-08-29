package user

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserService
}

func (h *UserHandler) InitDependencies() {
	repo := &UserRepository{}
	h.service = &UserService{
		repo: repo,
	}
}

func (h *UserHandler) Create(ctx *gin.Context) {
	h.service.Create(
		ctx,
	)
}

func (h *UserHandler) FindAll(ctx *gin.Context) {
	h.service.FindAll(
		ctx,
	)
}

func (h *UserHandler) FindOne(ctx *gin.Context) {
	h.service.FindOne(
		ctx,
	)
}

func (h *UserHandler) Update(ctx *gin.Context) {
	h.service.Update(
		ctx,
	)
}

func (h *UserHandler) Delete(ctx *gin.Context) {
	h.service.Delete(
		ctx,
	)
}
