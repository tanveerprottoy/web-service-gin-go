package user

import (
	"errors"
	"net/http"
	"txp/web-service-gin/src/module/user/dto"
	"txp/web-service-gin/src/module/user/entity"
	"txp/web-service-gin/src/util"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	repo *UserRepository
}

func (u *UserService) CreateUser(ctx *gin.Context) interface{} {
	var lastId int
	var p *dto.CreateUpdateUserBody
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		ctx.AbortWithError(
			http.StatusBadRequest,
			errors.New(
				util.BadRequest,
			),
		)
		return -1
	}
	lastId = u.repo.Create(
		p,
	)
	return lastId
}

func (u *UserService) FindUsers(ctx *gin.Context) []entity.User {
	users := u.repo.FindAll()
	return users
}

func (u *UserService) FindUser(ctx *gin.Context) entity.User {
	id, exists := ctx.Params.Get("id")
	if !exists {
		ctx.AbortWithError(
			http.StatusBadRequest,
			errors.New(
				util.BadRequest,
			),
		)
		return entity.User{}
	}
	user, err := u.repo.FindOne(
		id,
	)
	if err != nil {
		ctx.AbortWithError(
			http.StatusInternalServerError,
			errors.New(
				util.InternalServerError,
			),
		)
	}
	return user
}

func (u *UserService) UpdateUser(ctx *gin.Context) interface{} {
	id, exists := ctx.Params.Get("id")
	if !exists {
		ctx.AbortWithError(
			http.StatusBadRequest,
			errors.New(
				util.BadRequest,
			),
		)
		return -1
	}
	var p *entity.User
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		ctx.AbortWithError(
			http.StatusBadRequest,
			errors.New(
				util.BadRequest,
			),
		)
		return -1
	}
	rows := u.repo.Update(
		id,
		p,
	)
	return map[string]int64{util.RowsAffected: rows}
}
