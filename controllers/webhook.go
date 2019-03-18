package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"bitbucket.org/magazine-ondemand/adminplace-api/models"
	"bitbucket.org/magazine-ondemand/adminplace-api/repository"

	"github.com/gorilla/mux"
)

// GetWebhook valida a rota para receber chamada do workplace
func GetWebhook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json; charset=utf-8")
	verifyToken := r.URL.Query().Get("hub.verify_token")

	vars := mux.Vars(r)
	idIntegration, _ := strconv.Atoi(vars["id"])
	integration, _ := repository.GetIntegrationByID(idIntegration)

	w.WriteHeader(http.StatusOK)

	if verifyToken == integration.Verify {
		json.NewEncoder(w).Encode(r.URL.Query().Get("hub.challenge"))
	} else {
		json.NewEncoder(w).Encode("Error, wrong token")
	}
}

// PostWebhook recebe os eventos de mensagens
func PostWebhook(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	vars := mux.Vars(r)
	idIntegration, _ := strconv.Atoi(vars["id"])
	obj := decoderRequest(r, &models.MessagesWebhook{})
	mw := obj.(models.MessagesWebhook)

	messagingEvents := mw.Entry[0].Messaging
	var err error
	var up int64

	for _, m := range messagingEvents {
		log.Println(m)

		if len(m.Delivery.Mids) > 0 {
			up, err = repository.UpdateReceivedMessage(m.Sender.ID, idIntegration)
		}

		if m.Read.Watermark != 0 {
			up, err = repository.UpdateReadedMessage(m.Sender.ID, idIntegration)
		}
	}

	responseRequest(w, up, err)
}
