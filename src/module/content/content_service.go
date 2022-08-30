package content

import (
	"database/sql"
	"errors"
	"net/http"
	"txp/web-service-gin/src/module/content/dto"
	"txp/web-service-gin/src/module/content/entity"
	"txp/web-service-gin/src/util"

	"github.com/gin-gonic/gin"
)

type ContentService struct {
	repo *ContentRepository
}

func (s *ContentService) Create(ctx *gin.Context) {
	var b *dto.CreateUpdateContentBody
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

func (s *ContentService) FindAll(ctx *gin.Context) {
	Contents, err := s.repo.FindAll()
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
		Contents,
		ctx,
	)
}

func (s *ContentService) FindOne(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			ctx,
		)
		return
	}
	Content, err := s.repo.FindOne(
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
		Content,
		ctx,
	)
}

func (s *ContentService) Update(ctx *gin.Context) {
	id, exists := ctx.Params.Get("id")
	if !exists {
		util.ErrorAbort(
			http.StatusBadRequest,
			util.BadRequest,
			ctx,
		)
		return
	}
	var b *entity.Content
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

func (s *ContentService) Delete(ctx *gin.Context) {
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
