package handler

import (
	"biyelap.com/biyelap-core/app/datum/model"
	"biyelap.com/biyelap-core/app/core"
	"encoding/json"
	"net/http"
	"strconv"
)

type AuthHandler struct {
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// code for login
	var loginBody *model.LoginBody
	err := json.NewDecoder(r.Body).Decode(&loginBody)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	u, err := model.GetUserByPhone(loginBody.Phone)
	if err == nil {
		core.Respond(
			w,
			&model.LoginResponse{
				Id:           u.Id,
				IsActive:     u.IsActive,
				IsRegistered: u.IsRegistered,
				Token:        "",
			},
		)
		return
	}
	lastId := model.InsertUser(loginBody)
	var token = ""
	if lastId > 0 {
		token = core.GenerateToken(strconv.Itoa(lastId))
	}
	core.Respond(
		w,
		&model.LoginResponse{
			Id:           lastId,
			IsActive:     true,
			IsRegistered: false,
			Token:        token,
		},
	)
}

func (a *AuthHandler) PassAuth(w http.ResponseWriter, r *http.Request) {
	// code for pass PassAuth
	var p *model.PassAuthBody
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	h, err := model.GetUserPassHash(p.Id)
	var token = ""
	if err == nil {
		if core.CompareHashAndPassword(h, p.Pass) {
			token = core.GenerateToken(strconv.Itoa(p.Id))
		}
	}
	core.Respond(
		w,
		&model.PassAuthResponse{
			Token: token,
		},
	)
}

func (a *AuthHandler) RefreshToken(writer http.ResponseWriter, request *http.Request) {
	// block for RefreshToken
	core.Respond(writer, "")
}
