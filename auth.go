package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stonoy/start_mail/internal/database"
)

type My_Claims struct {
	role string
	jwt.RegisteredClaims
}

func generateToken(jetSecret string, user database.User) (string, error) {
	// create my claims
	claims := My_Claims{
		string(user.Role),
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "start_email",
			ID:        fmt.Sprintf("%v", user.ID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(jetSecret))

	return ss, err
}

func decodeToken(jetSecret, tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &My_Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jetSecret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*My_Claims); ok && token.Valid {
		return claims.ID, nil
	} else {
		return "", fmt.Errorf("token is not valid")
	}
}

func getTokenFromHeader(r *http.Request) (string, error) {
	// get auth header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("np auth header provided")
	}

	authHeaderSlice := strings.Fields(authHeader)

	if authHeaderSlice[0] == "Bearer" && len(authHeaderSlice) > 1 {
		return authHeaderSlice[1], nil
	} else {
		return "", errors.New("provide a valid header")
	}
}

type theAuthFuncType func(w http.ResponseWriter, r *http.Request, user database.User)

func (cfg *apiConfig) checkUserMiddleware(theFunc theAuthFuncType) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// get token from header
		token, err := getTokenFromHeader(r)
		if err != nil {
			respWithError(w, 401, fmt.Sprintf("%v", err))
			return
		}

		// decode the token
		useridStr, err := decodeToken(cfg.jwtSecret, token)
		if err != nil {
			respWithError(w, 401, fmt.Sprintf("%v", err))
			return
		}

		// parse the user id taken from token
		userId, err := uuid.Parse(useridStr)
		if err != nil {
			respWithError(w, 401, fmt.Sprintf("error in parsing user id -> %v", err))
			return
		}

		// get the user
		user, err := cfg.db.GetUserById(r.Context(), userId)
		if err != nil {
			if err == sql.ErrNoRows {
				respWithError(w, 400, fmt.Sprintf("no user with id -> %v", useridStr))
				return
			} else {
				respWithError(w, 500, fmt.Sprintf("error in GetUserById -> %v", err))
				return
			}
		}

		// call theFunc
		theFunc(w, r, user)
	}
}
