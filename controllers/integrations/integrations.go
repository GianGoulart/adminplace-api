package integrations

import (
	"fmt"
	"net/http"
	"strconv"

	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/utils"
	"bitbucket.org/magazine-ondemand/adminplace-api/models"
	"bitbucket.org/magazine-ondemand/adminplace-api/repository"

	"github.com/gorilla/mux"
)

// GetIntegrationByID rota: /integration/{id}
func GetIntegrationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	integration, err := repository.GetIntegrationByID(id)
	utils.ResponseRequest(w, integration, err)
}

// GetIntegrationByAny rota: /integration/{id}
func GetIntegrationByAny(w http.ResponseWriter, r *http.Request) {
	utils.ValidationRequest(w, r)
	obj := utils.DecoderRequest(r, &models.Integration{})
	i := obj.(*models.Integration)
	fmt.Println(i)
	integration, err := repository.GetIntegrationByAny(i)
	utils.ResponseRequest(w, integration, err)
}

// GetAllIntegration rota: /integration
func GetAllIntegration(w http.ResponseWriter, r *http.Request) {
	integration, err := repository.GetAllIntegration()
	utils.ResponseRequest(w, integration, err)
}

// CreateIntegration rota: /integration
func CreateIntegration(w http.ResponseWriter, r *http.Request) {
	utils.ValidationRequest(w, r)
	obj := utils.DecoderRequest(r, &models.Integration{})
	i := obj.(*models.Integration)
	integration, err := repository.CreateIntegration(i)
	utils.ResponseRequest(w, integration, err)
}

// UpdateIntegration rota: /integration
func UpdateIntegration(w http.ResponseWriter, r *http.Request) {
	utils.ValidationRequest(w, r)
	obj := utils.DecoderRequest(r, &models.Integration{})
	i := obj.(*models.Integration)
	fmt.Println(i)
	integration, err := repository.UpdateIntegration(i)
	utils.ResponseRequest(w, integration, err)
}

// DeleteIntegration rota: /integration/:id
func DeleteIntegration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	integration, err := repository.DeleteIntegration(id)
	utils.ResponseRequest(w, integration, err)
}
