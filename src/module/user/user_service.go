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

func (s *UserService) Create(ctx *gin.Context) {
	var b *dto.CreateUpdateUserBody
	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			ctx,
		)
		return
	}
	_, err = s.repo.Create(
		b,
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

func (s *UserService) FindAll(ctx *gin.Context) {
	users, err := s.repo.FindAll()
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

func (s *UserService) FindOne(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			ctx,
		)
		return
	}
	user, err := s.repo.FindOne(
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

func (s *UserService) Update(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			ctx,
		)
		return
	}
	var b *entity.User
	err := ctx.ShouldBindJSON(&b)
	if err != nil {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			ctx,
		)
		return
	}
	rows, err := s.repo.Update(
		id,
		b,
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

func (s *UserService) Delete(ctx *gin.Context) {
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
	rows, err := s.repo.Delete(
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
