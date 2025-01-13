package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/stonoy/start_mail/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
}

type Email struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	Sender    uuid.UUID `json:"sender"`
	Reciever  uuid.UUID `json:"receiver"`
}

type Favourite struct {
	ID          uuid.UUID `json:"id"`
	Userid      uuid.UUID `json:"userId"`
	CreatedAt   time.Time `json:"created_at"`
	ID_2        uuid.UUID `json:"emailId"`
	CreatedAt_2 time.Time `json:"email_created_at"`
	UpdatedAt   time.Time `json:"email_updated_at"`
	Subject     string    `json:"subject"`
	Body        string    `json:"body"`
	Sender      uuid.UUID `json:"sender"`
	Reciever    uuid.UUID `json:"reciver"`
}

func dbToRespEmail(emails []database.Email) []Email {
	final := []Email{}

	for _, email := range emails {
		final = append(final, Email{
			ID:        email.ID,
			CreatedAt: email.CreatedAt,
			UpdatedAt: email.UpdatedAt,
			Subject:   email.Subject,
			Body:      email.Body,
			Sender:    email.Sender,
			Reciever:  email.Reciever,
		})
	}

	return final
}

func dbToRespFav(favs []database.GetAllFavOfUserRow) []Favourite {
	final := []Favourite{}

	for _, fav := range favs {
		final = append(final, Favourite{
			ID:          fav.ID,
			CreatedAt:   fav.CreatedAt,
			UpdatedAt:   fav.UpdatedAt,
			Userid:      fav.Userid,
			CreatedAt_2: fav.CreatedAt_2,
			Subject:     fav.Subject,
			Body:        fav.Body,
			Sender:      fav.Sender,
			Reciever:    fav.Reciever,
			ID_2:        fav.ID_2,
		})
	}

	return final
}
