package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"txp/web-service-gin/src/core"
	"txp/web-service-gin/src/util"

	"github.com/IBM/ibm-cos-sdk-go/private/util"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *UserService
}

func (u *UserHandler) GetUsers(c *gin.Context) {
	s, err, p := u.userService.GetUsers(
		c,
	)
	if err != nil {
		util.RespondError(
			s,
			err,
			c,
		)
	}
	util.Respond(
		s,
		p,
		c,
	)
}

func (u *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, core.UrlKeyId)
	ownIdStr := r.URL.Query().Get("own_id")
	isAllColStr := r.URL.Query().Get("is_all_col")
	/*ctx := request.Context()
	key := ctx.Value(core.ContextKey).(string)
	fmt.Println(fmt.Sprintf("Context key %s", key))*/
	id, err := strconv.Atoi(userId)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	isAllCol, err := strconv.ParseBool(isAllColStr)
	if err != nil {
		isAllCol = false
	}
	user, err := model.GetUser(id, isAllCol)
	if err != nil {
		core.RespondError(w, err, http.StatusInternalServerError)
		return
	}
	isFav := false
	isConnected := false
	isRequested := false
	if ownIdStr != "" {
		ownId, err := strconv.Atoi(ownIdStr)
		if err != nil {
			core.RespondError(w, err, http.StatusBadRequest)
			return
		}
		f := model.GetFavouriteRowCount(id, ownId)
		if f > 0 {
			isFav = true
		}
		c := model.GetConnectionRowCount(id, ownId)
		if c > 0 {
			isConnected = true
		}
		if !isConnected {
			r := model.GetConnectionRequestCount(id, ownId)
			if r > 0 {
				isRequested = true
			}
		}
	}
	core.Respond(
		w,
		map[string]interface{}{
			"user":        user,
			"isFav":       isFav,
			"isConnected": isConnected,
			"isRequested": isRequested,
		},
	)
}

func (u *UserHandler) PutUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, core.UrlKeyId)
	id, err := strconv.Atoi(userId)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	var user *model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		core.RespondError(w, err, http.StatusUnauthorized)
		return
	}
	fmt.Print(user)
	//	ctx := request.Context()
	//	key := ctx.Value("key").(string)
	//	fmt.Print(key)
	user.Id = id
	rows := model.UpdateUser(user)
	core.Respond(w, map[string]int64{core.RowsAffected: rows})
}
