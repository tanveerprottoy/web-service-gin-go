package user

import (
	"errors"
	"net/http"
	"txp/web-service-gin/src/util"

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
	p := u.service.CreateUser(
		ctx,
	)
	if p == -1 {
		util.RespondError(
			http.StatusBadRequest,
			errors.New(
				util.BadRequest,
			),
			ctx,
		)
		return
	}
	util.Respond(
		http.StatusCreated,
		map[string]bool{
			"created": true,
		},
		ctx,
	)
}

func (u *UserHandler) FindUsers(ctx *gin.Context) {
	p := u.service.FindUsers(
		ctx,
	)
	util.Respond(
		http.StatusOK,
		p,
		ctx,
	)
}

func (u *UserHandler) FindUser(ctx *gin.Context) {
	p := u.service.FindUsers(
		ctx,
	)
	util.Respond(
		http.StatusOK,
		p,
		ctx,
	)
}

func (u *UserHandler) UpdateUser(ctx *gin.Context) {
	p := u.service.UpdateUser(
		ctx,
	)
	util.Respond(
		http.StatusOK,
		p,
		ctx,
	)
}
