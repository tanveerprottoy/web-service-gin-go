package user

import (
	"net/http"
	"strconv"
	"txp/web-service-gin/src/modules/user/entity"
	"txp/web-service-gin/src/util"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	
}

func (u *UserService) GetUsers(c *gin.Context) (int, error, []entity.User) {
	// for seek/key-set pagination
	lastIdStr := c.Query(util.LastId)
	lastId, err := strconv.Atoi(lastIdStr)
	if err != nil {
		return http.StatusBadRequest, err, []entity.User{}
	}
	users := UserRepositoy.GetUsers(
		id,
		lastId,
		gender,
		religion,
		maritalStatus,
	)
	return users
}