package controllers

import (
	"api/security"
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	responseBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(responseBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	savedUserDB, err := repository.FindByEmail(user.Email)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(savedUserDB.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.CreateToken(savedUserDB.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(token)
}
