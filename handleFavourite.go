package main

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stonoy/start_mail/internal/database"
)

func (cfg *apiConfig) createFav(w http.ResponseWriter, r *http.Request, user database.User) {
	emailIdStr := chi.URLParam(r, "emailID")

	emailId, err := uuid.Parse(emailIdStr)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in parsing str to uuid -> %v", err))
		return
	}

	fav, err := cfg.db.AddToFavourites(r.Context(), database.AddToFavouritesParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Userid:    user.ID,
		Emailid:   emailId,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in AddToFavourites -> %v", err))
		return
	}

	respWithJson(w, 201, struct {
		Msg string `json:"msg"`
	}{
		Msg: fmt.Sprintf("fav created -> %v", fav.Emailid),
	})
}

func (cfg *apiConfig) getAllFav(w http.ResponseWriter, r *http.Request, user database.User) {
	// get the query
	queryParams := r.URL.Query()

	pageQ := queryParams.Get("page")
	page := 1

	if pageQ != "" {
		pageInt, err := strconv.Atoi(pageQ)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("error in converting page %v", err))
			return
		}

		page = pageInt
	}

	limit := 2
	offset := (page - 1) * limit

	favs, err := cfg.db.GetAllFavOfUser(r.Context(), database.GetAllFavOfUserParams{
		Userid: user.ID,
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 200, fmt.Sprintf("no fav email with user -> %v", user.Name))
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in GetAllFavOfUser -> %v", err))
			return
		}
	}

	numOfFavUser, err := cfg.db.NumOfAllFavUser(r.Context(), user.ID)
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in NumOfAllFavUser -> %v", err))
		return
	}

	numOfPages := math.Ceil(float64(numOfFavUser) / float64(limit))

	// response
	type respStruct struct {
		Favourite  []Favourite `json:"favourite"`
		NumOfPages float64     `json:"numOfPages"`
		Page       int         `json:"page"`
	}

	respWithJson(w, 200, respStruct{
		Favourite:  dbToRespFav(favs),
		NumOfPages: numOfPages,
		Page:       page,
	})

}

func (cfg *apiConfig) deleteFav(w http.ResponseWriter, r *http.Request, user database.User) {
	favIdStr := chi.URLParam(r, "favID")

	favId, err := uuid.Parse(favIdStr)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in parsing str to uuid -> %v", err))
		return
	}

	deletedFav, err := cfg.db.DeleteFav(r.Context(), database.DeleteFavParams{
		ID:     favId,
		Userid: user.ID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, fmt.Sprintf("no fav email to delete with user, fav id -> %v", favId))
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in DeleteFav -> %v", err))
			return
		}
	}

	respWithJson(w, 200, struct {
		Msg string `json:"msg"`
	}{
		Msg: fmt.Sprintf("fav deleted -> %v", deletedFav.ID),
	})
}
