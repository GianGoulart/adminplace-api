package controllers

import (
	"net/http"
	"strconv"

	"bitbucket.org/dt_souza/adminplace-api/models"
	"bitbucket.org/dt_souza/adminplace-api/repository"
	"github.com/gorilla/mux"
)

// GetWelcomeByID rota: /welcome/{id}
func GetWelcomeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	welcome, err := repository.GetWelcomeByID(id)
	responseRequest(w, welcome, err)
}

// GetAllWelcome rota: /welcome
func GetAllWelcome(w http.ResponseWriter, r *http.Request) {
	welcome, err := repository.GetAllWelcome()
	responseRequest(w, welcome, err)
}

// CreateWelcome rota: /welcome
func CreateWelcome(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	obj := decoderRequest(r, &models.Welcome{})
	us := obj.(models.Welcome)

	welcome, err := repository.CreateWelcome(us)
	responseRequest(w, welcome, err)
}

// UpdateWelcome rota: /welcome
func UpdateWelcome(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	obj := decoderRequest(r, &models.Welcome{})
	us := obj.(models.Welcome)

	welcome, err := repository.UpdateWelcome(us)
	responseRequest(w, welcome, err)
}

// DeleteWelcome rota: /welcome/:id
func DeleteWelcome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	welcome, err := repository.DeleteWelcome(id)
	responseRequest(w, welcome, err)
}
