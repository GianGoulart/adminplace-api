package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"bitbucket.org/dt_souza/adminplace-api/repository"
	"bitbucket.org/dt_souza/auditoria-remota/models"
)

// GetWebhook valida a rota para receber chamada do workplace
func GetWebhook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/json; charset=utf-8")
	verifyToken := r.URL.Query().Get("hub.verify_token")
	integration, _ := repository.GetIntegrationByVerify(verifyToken)

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
	obj := decoderRequest(r, &models.MessagesWebhook{})
	mw := obj.(models.MessagesWebhook)

	messagingEvents := mw.Entry[0].Messaging
	var err error
	var up int64

	for _, m := range messagingEvents {
		log.Println(m)

		if len(m.Delivery.Mids) > 0 {
			up, err = repository.UpdateReceivedMessage(m.Sender.ID)
		}

		if m.Read.Watermark != 0 {
			up, err = repository.UpdateReadedMessage(m.Sender.ID)
		}
	}

	responseRequest(w, up, err)
}
