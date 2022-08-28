package user

import (
	"database/sql"
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

func (u *UserService) CreateUser(ctx *gin.Context) {
	var p *dto.CreateUpdateUserBody
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			ctx,
		)
		return
	}
	_, err = u.repo.Create(
		p,
	)
	if err != nil {
		util.ErrorAbort(
			http.StatusInternalServerError,
			err.Error(),
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

func (u *UserService) FindUsers(ctx *gin.Context) {
	users, err := u.repo.FindAll()
	if err != nil {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			ctx,
		)
		return
	}
	util.Respond(
		http.StatusOK,
		users,
		ctx,
	)
}

func (u *UserService) FindUser(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			ctx,
		)
		return
	}
	user, err := u.repo.FindOne(
		id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			util.ErrorAbort(
				http.StatusNotFound,
				util.NotFound,
				ctx,
			)
			return
		}
		ctx.AbortWithError(
			http.StatusInternalServerError,
			errors.New(
				util.InternalServerError,
			),
		)
		/* util.ErrorAbort(
			http.StatusInternalServerError,
			util.InternalServerError,
			ctx,
		) */
		return
	}
	util.Respond(
		http.StatusOK,
		user,
		ctx,
	)
}

func (u *UserService) UpdateUser(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			ctx,
		)
		return
	}
	var p *entity.User
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			ctx,
		)
		return
	}
	rows, err := u.repo.Update(
		id,
		p,
	)
	if err != nil {
		util.ErrorAbort(
			http.StatusInternalServerError,
			util.InternalServerError,
			ctx,
		)
		return
	}
	if rows > 0 {
		util.Respond(
			http.StatusOK,
			map[string]int64{util.RowsAffected: rows},
			ctx,
		)
		return
	}
	util.ErrorAbort(
		http.StatusInternalServerError,
		util.InternalServerError,
		ctx,
	)
}

func (u *UserService) DeleteUser(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		util.RespondError(
			http.StatusBadRequest,
			errors.New(
				util.BadRequest,
			),
			ctx,
		)
		return
	}
	rows, err := u.repo.Delete(
		id,
	)
	if err != nil {
		util.ErrorAbort(
			http.StatusInternalServerError,
			util.InternalServerError,
			ctx,
		)
		return
	}
	util.Respond(
		http.StatusOK,
		map[string]int64{util.RowsAffected: rows},
		ctx,
	)
}
