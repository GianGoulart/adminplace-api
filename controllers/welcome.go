package controllers

import (
	"fmt"
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

// SendWelcomeMessage envia a mensagem de boas vindas para novos colaboradores
func SendWelcomeMessage() {
	empl, err := repository.GetEmployeeByWelcome(false)
	if err != nil {
		fmt.Println(err)
	}

	msg, err := repository.GetWelcomeByActive(true)
	if err != nil {
		fmt.Println(err)
	}

	for _, e := range empl {
		if e.IDWorkplace == "" {
			wpUser, err := buscaWorkplaceUser(e.Email)
			if err == nil {
				_, err = sendTextMessage(wpUser.ID, msg.Text)
				if err == nil {
					e.Welcome = true
					repository.UpdateEmployee(e)
				}
			}
		} else {
			_, err = sendTextMessage(e.IDWorkplace, msg.Text)
			if err == nil {
				e.Welcome = true
				repository.UpdateEmployee(e)
			}
		}
	}
}
