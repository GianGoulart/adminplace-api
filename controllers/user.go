package controllers

import (
	"net/http"
	"strconv"

	"bitbucket.org/dt_souza/adminplace-api/models"
	"bitbucket.org/dt_souza/adminplace-api/repository"
	"github.com/gorilla/mux"
)

// GetUserByID rota: /user/{id}
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	user, err := repository.GetUserByID(id)
	responseRequest(w, user, err)
}

// GetAllUser rota: /user
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	user, err := repository.GetAllUser()
	responseRequest(w, user, err)
}

// CreateUser rota: /user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	obj := decoderRequest(r, &models.User{})
	us := obj.(models.User)

	user, err := repository.CreateUser(us)
	responseRequest(w, user, err)
}

// UpdateUser rota: /user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	obj := decoderRequest(r, &models.User{})
	us := obj.(models.User)

	user, err := repository.UpdateUser(us)
	responseRequest(w, user, err)
}

// DeleteUser rota: /user/:id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	user, err := repository.DeleteUser(id)
	responseRequest(w, user, err)
}
