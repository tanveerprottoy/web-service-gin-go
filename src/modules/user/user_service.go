package user

import (
	"net/http"
	"strconv"
	"txp/web-service-gin/src/modules/user/entity"
	"txp/web-service-gin/src/util"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepositoy := UsereUserRepositoy{}
}

func (u *UserService) GetUsers(ctx *gin.Context) (int, error, []entity.User) {
	// for seek/key-set pagination
	lastIdStr := ctx.Query(util.LastId)
	lastId, err := strconv.Atoi(lastIdStr)
	if err != nil {
		return http.StatusBadRequest, err, []entity.User{}
	}
	users := UserRepositoy.GetUsers()
	return users
}