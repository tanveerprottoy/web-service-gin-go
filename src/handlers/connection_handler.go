package handler

import (
	"net/http"
	"strconv"

	"biyelap.com/biyelap-core/app/core"
	"biyelap.com/biyelap-core/app/datum/model"
	"github.com/go-chi/chi"
)

type ConnectionHandler struct {
}

func (c *ConnectionHandler) GetConnections(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, core.UrlKeyId)
	userId, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	connections := model.GetConnections(userId)
	core.Respond(w, connections)
}

/* func (c *ConnectionHandler) PostConnection(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, core.UrlKeyId)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	var conn *model.ConnectionUser
	err = json.NewDecoder(r.Body).Decode(&conn)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	lastId := model.InsertConnection(id, conn)
	core.Respond(w, map[string]int{core.LastId: lastId})
} */

func (c *ConnectionHandler) DeleteConnection(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, core.UrlKeyId)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	connIdStr := chi.URLParam(r, core.UrlKeyConn)
	connId, err := strconv.Atoi(connIdStr)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	rows := model.DeleteConnection(id, connId)
	core.Respond(w, map[string]int64{core.RowsAffected: rows})
}
