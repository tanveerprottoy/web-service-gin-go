package handler

import (
	"biyelap.com/biyelap-core/app/datum/model"
	"biyelap.com/biyelap-core/app/core"
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type FavouriteHandler struct {
}

func (f *FavouriteHandler) GetFavourites(writer http.ResponseWriter, request *http.Request) {
	idParam := chi.URLParam(request, core.UrlKeyId)
	userId, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondError(writer, err, http.StatusUnauthorized)
		return
	}
	favourites := model.GetFavourites(userId)
	core.Respond(writer, favourites)
}

func (f *FavouriteHandler) PostFavourite(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, core.UrlKeyId)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	var fav *model.FavouriteUser
	err = json.NewDecoder(r.Body).Decode(&fav)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	lastId := model.InsertFavourite(id, fav)
	core.Respond(w, map[string]int{core.LastId: lastId})
}

func (f *FavouriteHandler) DeleteFavourite(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, core.UrlKeyId)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	favIdStr := chi.URLParam(r, core.UrlKeyFav)
	favId, err := strconv.Atoi(favIdStr)
	if err != nil {
		core.RespondError(w, err, http.StatusBadRequest)
		return
	}
	rows := model.DeleteFavourite(id, favId)
	core.Respond(w, map[string]int64{core.RowsAffected: rows})
}
