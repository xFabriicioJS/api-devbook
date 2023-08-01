package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser is responsible for creating a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var usuario models.User

	if err = json.Unmarshal(requestBody, &usuario); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUsersRepository(db)
	repository.Create(usuario)
}

// FindAllUsers is responsible for getting all users
func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all users"))
}

// FindUser is responsible for getting a user
func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Finding user..."))
}

// UpdateUser is responsible for updating a user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user..."))
}

// DeleteUser is responsible for deleting a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user..."))
}
