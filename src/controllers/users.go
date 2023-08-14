package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser is responsible for creating a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("signup"); err != nil {
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
	user.ID, err = repository.Create(user)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// FindAllUsers is responsible for getting all users
func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, err := database.Connect()

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	users, err := repository.Find(nameOrNick)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// FindUser is responsible for getting a user
func FindUser(w http.ResponseWriter, r *http.Request) {

	// Getting all parameters in route
	parameters := mux.Vars(r)

	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)

	if err != nil {
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

	user, err := repository.FindById(userId)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser is responsible for updating a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, err := strconv.ParseUint(parameters["usersId"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	userIdToken, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("It's impossible to edit a user that isn't yours."))
		return
	}

	if userIdToken != userId {
		responses.Error(w, http.StatusForbidden, err)
		return
	}

	responseBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Error(
			w,
			http.StatusUnprocessableEntity,
			err,
		)
		return
	}

	var user models.User

	if err = json.Unmarshal(responseBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare("update"); err != nil {
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
	if err = repository.Update(userId, user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser is responsible for deleting a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	userIdToken, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userIdToken != userId {
		responses.Error(w, http.StatusForbidden, errors.New("It's impossible to delete a user different from yours"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)

	if err = repository.Delete(userId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := auth.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	userId, err := strconv.ParseUint(parameters["userId"], 10, 64)

	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if followerId == userId {
		responses.Error(w, http.StatusForbidden, errors.New("It's impossible to follow yourself"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	if err = repository.Follow(userId, followerId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
}
