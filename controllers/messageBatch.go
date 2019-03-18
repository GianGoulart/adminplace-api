package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"bitbucket.org/magazine-ondemand/adminplace-api/models"
	"bitbucket.org/magazine-ondemand/adminplace-api/repository"
	"github.com/gorilla/mux"
)

// GetMessageBatchByID rota: /batch/{id}
func GetMessageBatchByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	batch, err := repository.GetMessageBatchByID(id)
	responseRequest(w, batch, err)
}

// CreateMessageBatch rota: /batch
func CreateMessageBatch(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	obj := decoderRequest(r, &models.MessageBatch{})
	btc := obj.(models.MessageBatch)

	batch, err := repository.CreateMessageBatch(btc)
	responseRequest(w, batch, err)
}

// GetMessageBatchByAny rota: /batch/{id}/message
func GetMessageBatchByAny(w http.ResponseWriter, r *http.Request) {
	fmt.Println("vem")
	validationRequest(w, r)
	obj := decoderRequest(r, &models.MessageBatch{})
	btc := obj.(*models.MessageBatch)

	batch, err := repository.GetMessageBatchByAny(btc)
	responseRequest(w, batch, err)
}
