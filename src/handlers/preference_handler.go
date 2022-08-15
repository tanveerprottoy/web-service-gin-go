package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"biyelap.com/biyelap-core/app/core"
	"biyelap.com/biyelap-core/app/datum/model"
	"github.com/go-chi/chi"
)

type PreferenceHandler struct {
}

func (p *PreferenceHandler) GetPreference(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, core.UrlKeyId)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	pref := model.GetPreference(id)
	core.Respond(w, pref)
}

func (p *PreferenceHandler) PostPreference(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, core.UrlKeyId)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	var pref *model.Preference
	err = json.NewDecoder(r.Body).Decode(&pref)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	lastId := model.InsertPreference(pref, id)
	core.Respond(w, map[string]int{core.LastId: lastId})
}

func (p *PreferenceHandler) PutPreference(w http.ResponseWriter, r *http.Request) {
	preferenceIdStr := chi.URLParam(r, core.UrlKeyPreference)
	preferenceId, err := strconv.Atoi(preferenceIdStr)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	var pref *model.Preference
	err = json.NewDecoder(r.Body).Decode(&pref)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	rows := model.UpdatePreference(pref, preferenceId)
	core.Respond(w, map[string]int64{core.RowsAffected: rows})
}
