package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stonoy/start_mail/internal/database"
)

func (cfg *apiConfig) createEmail(w http.ResponseWriter, r *http.Request, user database.User) {
	// get the inputs
	type reqStruct struct {
		Recipient string `json:"recipient"`
		Subject   string `json:"subject"`
		Body      string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	reqObj := reqStruct{}
	err := decoder.Decode(&reqObj)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("can not decode request body -> %v", err))
		return
	}

	// check details
	if reqObj.Recipient == "" || reqObj.Subject == "" || reqObj.Body == "" {
		respWithError(w, 400, "provide correct credentials")
		return
	}

	// parse the user id taken from request body
	recipientId, err := uuid.Parse(reqObj.Recipient)
	if err != nil {
		respWithError(w, 401, fmt.Sprintf("error in parsing user id -> %v", err))
		return
	}

	// create email
	email, err := cfg.db.CreateMail(r.Context(), database.CreateMailParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Subject:   reqObj.Subject,
		Body:      reqObj.Body,
		Sender:    user.ID,
		Reciever:  recipientId,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in CreateMail -> %v", err))
		return
	}

	// send response
	respWithJson(w, 201, struct {
		Email Email `json:"email"`
	}{
		Email: Email{
			ID:        email.ID,
			CreatedAt: email.CreatedAt,
			UpdatedAt: email.UpdatedAt,
			Subject:   email.Subject,
			Body:      email.Body,
			Reciever:  email.Reciever,
			Sender:    email.Sender,
		},
	})
}

func (cfg *apiConfig) Inbox(w http.ResponseWriter, r *http.Request, user database.User) {
	// get the query
	queryParams := r.URL.Query()

	bodyQ := queryParams.Get("body")
	startTimeQ := queryParams.Get("start_time")
	endTimeQ := queryParams.Get("end_time")
	pageQ := queryParams.Get("page")

	// initiate query params
	body := "%%"
	startTime := time.Now().AddDate(-1, 0, 0)
	endTime := time.Now().AddDate(1, 0, 0)
	page := 1

	if bodyQ != "" {
		body = "%" + bodyQ + "%"
	}

	if pageQ != "" {
		pageInt, err := strconv.Atoi(pageQ)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("error in converting page %v", err))
			return
		}

		page = pageInt
	}

	if startTimeQ != "" {
		theTime, err := getTimeFromStr(startTimeQ)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("%v", err))
			return
		}

		startTime = theTime
	}

	if endTimeQ != "" {
		theTime, err := getTimeFromStr(endTimeQ)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("%v", err))
			return
		}

		endTime = theTime
	}

	// set limit and offset
	limit := 2
	offset := (page - 1) * limit

	// get the mails with filter
	emails, err := cfg.db.InboxMailWithFilter(r.Context(), database.InboxMailWithFilterParams{
		Reciever:    user.ID,
		Body:        body,
		CreatedAt:   startTime,
		CreatedAt_2: endTime,
		Limit:       int32(limit),
		Offset:      int32(offset),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 200, fmt.Sprintf("no email in inbox with user -> %v", user.Name))
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in InboxMailWithFilter -> %v", err))
			return
		}
	}

	// get the numofmails
	numOfEmailsWithFilter, err := cfg.db.NumOfInboxMailWithFilter(r.Context(), database.NumOfInboxMailWithFilterParams{
		Reciever:    user.ID,
		Body:        body,
		CreatedAt:   startTime,
		CreatedAt_2: endTime,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in NumOfInboxMailWithFilter -> %v", err))
		return
	}

	numOfPages := math.Ceil(float64(numOfEmailsWithFilter) / float64(limit))

	// response
	type respStruct struct {
		Emails     []Email `json:"emails"`
		NumOfPages float64 `json:"numOfPages"`
		Page       int     `json:"page"`
	}

	respWithJson(w, 200, respStruct{
		Emails:     dbToRespEmail(emails),
		NumOfPages: numOfPages,
		Page:       page,
	})
}

func (cfg *apiConfig) SentBox(w http.ResponseWriter, r *http.Request, user database.User) {
	// get the query
	queryParams := r.URL.Query()

	bodyQ := queryParams.Get("body")
	startTimeQ := queryParams.Get("start_time")
	endTimeQ := queryParams.Get("end_time")
	pageQ := queryParams.Get("page")

	// initiate query params
	body := "%%"
	startTime := time.Now().AddDate(-1, 0, 0)
	endTime := time.Now().AddDate(1, 0, 0)
	page := 1

	if bodyQ != "" {
		body = "%" + bodyQ + "%"
	}

	if pageQ != "" {
		pageInt, err := strconv.Atoi(pageQ)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("error in converting page %v", err))
			return
		}

		page = pageInt
	}

	if startTimeQ != "" {
		theTime, err := getTimeFromStr(startTimeQ)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("%v", err))
			return
		}

		startTime = theTime
	}

	if endTimeQ != "" {
		theTime, err := getTimeFromStr(endTimeQ)
		if err != nil {
			respWithError(w, 400, fmt.Sprintf("%v", err))
			return
		}

		endTime = theTime
	}

	// set limit and offset
	limit := 2
	offset := (page - 1) * limit

	// get the mails with filter
	emails, err := cfg.db.SentBoxMailWithFilter(r.Context(), database.SentBoxMailWithFilterParams{
		Sender:      user.ID,
		Body:        body,
		CreatedAt:   startTime,
		CreatedAt_2: endTime,
		Limit:       int32(limit),
		Offset:      int32(offset),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 200, fmt.Sprintf("no email in sent inbox with user -> %v", user.Name))
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in SentBoxMailWithFilter -> %v", err))
			return
		}
	}

	// get the numofmails
	numOfEmailsWithFilter, err := cfg.db.NumOfSentBoxMailWithFilter(r.Context(), database.NumOfSentBoxMailWithFilterParams{
		Sender:      user.ID,
		Body:        body,
		CreatedAt:   startTime,
		CreatedAt_2: endTime,
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in NumOfSentBoxMailWithFilter -> %v", err))
		return
	}

	numOfPages := math.Ceil(float64(numOfEmailsWithFilter) / float64(limit))

	// response
	type respStruct struct {
		Emails     []Email `json:"emails"`
		NumOfPages float64 `json:"numOfPages"`
		Page       int     `json:"page"`
	}

	respWithJson(w, 200, respStruct{
		Emails:     dbToRespEmail(emails),
		NumOfPages: numOfPages,
		Page:       page,
	})
}

func (cfg *apiConfig) getSingleMail(w http.ResponseWriter, r *http.Request, user database.User) {
	emailIdStr := chi.URLParam(r, "emailID")

	emailId, err := uuid.Parse(emailIdStr)
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in parsing str to uuid -> %v", err))
		return
	}

	email, err := cfg.db.GetMailById(r.Context(), database.GetMailByIdParams{
		ID:       emailId,
		Sender:   user.ID,
		Reciever: user.ID,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, fmt.Sprintf("no such email found -> %v", emailIdStr))
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in GetMailById -> %v", err))
			return
		}
	}

	type respStruct struct {
		Email Email `json:"email"`
	}

	respWithJson(w, 200, respStruct{
		Email: Email{
			ID:        email.ID,
			CreatedAt: email.CreatedAt,
			UpdatedAt: email.UpdatedAt,
			Subject:   email.Subject,
			Body:      email.Body,
			Sender:    email.Sender,
			Reciever:  email.Reciever,
		},
	})

}
