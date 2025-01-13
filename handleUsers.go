package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/stonoy/start_mail/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) register(w http.ResponseWriter, r *http.Request) {
	// get user details
	type reqStruct struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	reqObj := reqStruct{}
	err := decoder.Decode(&reqObj)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("can not decode request body -> %v", err))
		return
	}

	// check details
	if reqObj.Name == "" || reqObj.Email == "" || reqObj.Password == "" || len(reqObj.Password) < 6 {
		respWithError(w, 400, "provide correct credentials")
		return
	}

	// hash password
	hashedPassowrd, err := bcrypt.GenerateFromPassword([]byte(reqObj.Password), 14)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("can not hash password -> %v", err))
		return
	}

	role := "user"

	// check for admin-user
	isAdmin, err := cfg.db.IsAdmin(r.Context())
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in IsAdmin -> %v", err))
		return
	}

	if isAdmin {
		role = "admin"
	}

	// create user
	user, err := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      reqObj.Name,
		Email:     reqObj.Email,
		Password:  string(hashedPassowrd),
		Role:      database.UserType(role),
	})
	if err != nil {
		respWithError(w, 500, fmt.Sprintf("error in CreateUser -> %v", err))
		return
	}

	// create token
	token, err := generateToken(cfg.jwtSecret, user)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in generateToken -> %v", err))
		return
	}

	// send response
	type respStruct struct {
		Token string `json:"token"`
		User  User   `json:"user"`
	}

	respWithJson(w, 201, respStruct{
		Token: token,
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Name:      user.Name,
			Email:     user.Email,
			Role:      string(user.Role),
		},
	})
}

func (cfg *apiConfig) login(w http.ResponseWriter, r *http.Request) {
	// get user details
	type reqStruct struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	reqObj := reqStruct{}
	err := decoder.Decode(&reqObj)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("can not decode request body -> %v", err))
		return
	}

	// check details
	if reqObj.Email == "" || reqObj.Password == "" || len(reqObj.Password) < 6 {
		respWithError(w, 400, "provide correct credentials")
		return
	}

	// get the user based on given email
	user, err := cfg.db.GetUserByEmail(r.Context(), reqObj.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			respWithError(w, 400, fmt.Sprintf("no user with email -> %v", reqObj.Email))
			return
		} else {
			respWithError(w, 500, fmt.Sprintf("error in GetUserByEmail -> %v", err))
			return
		}
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqObj.Password))
	if err != nil {
		respWithError(w, 401, "password not matched")
		return
	}

	// create token
	token, err := generateToken(cfg.jwtSecret, user)
	if err != nil {
		respWithError(w, 400, fmt.Sprintf("error in generateToken -> %v", err))
		return
	}

	// send response
	type respStruct struct {
		Token string `json:"token"`
		User  User   `json:"user"`
	}

	respWithJson(w, 201, respStruct{
		Token: token,
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Name:      user.Name,
			Email:     user.Email,
			Role:      string(user.Role),
		},
	})
}
