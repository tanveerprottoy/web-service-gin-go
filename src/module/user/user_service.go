package user

import (
	"errors"
	"net/http"
	"strconv"
	"txp/web-service-gin/src/module/user/entity"
	"txp/web-service-gin/src/util"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	repo *UserRepository
}

func (u *UserService) FindUsers(ctx *gin.Context) (
	[]entity.User,
	error,
) {
	// for seek/key-set pagination
	lastIdStr := ctx.Query(util.LastId)
	_, err := strconv.Atoi(lastIdStr)
	if err != nil {
		ctx.AbortWithError(
			http.StatusNotFound,
			errors.New(
				util.NotFound,
			),
		)
		return []entity.User{}, err
	}
	users := u.repo.FindAll()
	return users, nil
}
