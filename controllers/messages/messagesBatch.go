package messages

import (
	"fmt"
	"net/http"
	"strconv"

	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/utils"
	"bitbucket.org/magazine-ondemand/adminplace-api/models"
	"bitbucket.org/magazine-ondemand/adminplace-api/repository"

	"github.com/gorilla/mux"
)

// GetMessageBatchByID rota: /batch/{id}
func GetMessageBatchByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	batch, err := repository.GetMessageBatchByID(id)
	utils.ResponseRequest(w, batch, err)
}

// CreateMessageBatch rota: /batch
func CreateMessageBatch(w http.ResponseWriter, r *http.Request) {
	utils.ValidationRequest(w, r)
	obj := utils.DecoderRequest(r, &models.MessageBatch{})
	btc := obj.(models.MessageBatch)

	batch, err := repository.CreateMessageBatch(btc)
	utils.ResponseRequest(w, batch, err)
}

// GetMessageBatchByAny rota: /batch/{id}/message
func GetMessageBatchByAny(w http.ResponseWriter, r *http.Request) {
	fmt.Println("vem")
	utils.ValidationRequest(w, r)
	obj := utils.DecoderRequest(r, &models.MessageBatch{})
	btc := obj.(*models.MessageBatch)

	batch, err := repository.GetMessageBatchByAny(btc)
	utils.ResponseRequest(w, batch, err)
}
