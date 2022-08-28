package user

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserService
}

func (u *UserHandler) InitDependencies() {
	repo := &UserRepository{}
	u.service = &UserService{
		repo: repo,
	}
}

func (u *UserHandler) CreateUser(ctx *gin.Context) {
	u.service.CreateUser(
		ctx,
	)
}

func (u *UserHandler) FindUsers(ctx *gin.Context) {
	u.service.FindUsers(
		ctx,
	)
}

func (u *UserHandler) FindUser(ctx *gin.Context) {
	u.service.FindUser(
		ctx,
	)
}

func (u *UserHandler) UpdateUser(ctx *gin.Context) {
	u.service.UpdateUser(
		ctx,
	)
}

func (u *UserHandler) DeleteUser(ctx *gin.Context) {
	u.service.DeleteUser(
		ctx,
	)
}
