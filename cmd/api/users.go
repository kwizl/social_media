package main

import (
	"net/http"

	"github.com/kwizl/social_media/internal/store"
)

type CreateUserPayload struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var payload *CreateUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	user := &store.User{
		Firstname: payload.Firstname,
		Lastname: payload.Lastname,
		Email: payload.Email,
		Password: payload.Password,
		Username: payload.Username,
	}

	ctx := r.Context()
	if err := app.store.Users.Create(ctx, user); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
