package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"biyelap.com/biyelap-core/app/core"
	"biyelap.com/biyelap-core/app/datum/model"
	"github.com/go-chi/chi"
)

type ConnectionRequestHandler struct {
}

func (c *ConnectionRequestHandler) GetConnectionRequests(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, core.UrlKeyId)
	userId, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	connReqs := model.GetConnectionRequests(userId)
	core.Respond(w, connReqs)
}

func (c *ConnectionRequestHandler) PostConnectionRequest(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, core.UrlKeyId)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	var connReq *model.ConnectionRequest
	err = json.NewDecoder(r.Body).Decode(&connReq)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	lastId := model.InsertConnectionRequest(id, connReq)
	core.Respond(w, map[string]int{core.LastId: lastId})
}

func (c *ConnectionRequestHandler) PutConnectionRequest(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, core.UrlKeyId)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	connReqIDStr := chi.URLParam(r, core.UrlKeyConnReqId)
	connReqID, err := strconv.Atoi(connReqIDStr)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	var connReq *model.ConnectionRequest
	err = json.NewDecoder(r.Body).Decode(&connReq)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	rows := model.UpdateConnectionRequest(connReqID, connReq)
	if rows > 0 {
		conn := &model.ConnectionUser{}
		conn.ConnectionId = connReq.RequesterID
		lastId := model.InsertConnection(id, conn)
		if lastId < 1 {
			rows = 0
		}
	}
	core.Respond(w, map[string]int64{core.RowsAffected: rows})
}

func (c *ConnectionRequestHandler) DeleteConnectionRequest(w http.ResponseWriter, r *http.Request) {
	connReqIDStr := chi.URLParam(r, core.UrlKeyConnReqId)
	connReqID, err := strconv.Atoi(connReqIDStr)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	rows := model.DeleteConnectionRequest(connReqID)
	core.Respond(w, map[string]int64{core.RowsAffected: rows})
}
