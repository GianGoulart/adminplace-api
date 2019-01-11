package controllers

import (
	"net/http"
	"os/exec"
	"strconv"

	"bitbucket.org/dt_souza/adminplace-api/models"
	"bitbucket.org/dt_souza/adminplace-api/repository"
	"github.com/gorilla/mux"
)

// GetMessageByID rota: /message/{id}
func GetMessageByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	message, err := repository.GetMessageByID(id)
	responseRequest(w, message, err)
}

// GetMessageByBatch rota: /batch/{id}/message
func GetMessageByBatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	batch, err := repository.GetMessageByBatch(id)
	responseRequest(w, batch, err)
}

// CreateMessage rota: /message
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	obj := decoderRequest(r, &models.Message{})
	msg := obj.(models.Message)

	message, err := repository.CreateMessage(msg)
	responseRequest(w, message, err)
}

// UpdateReceivedMessage rota: /message/{id}/receive
func UpdateReceivedMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	message, err := repository.UpdateReceivedMessage(id)
	responseRequest(w, message, err)
}

// UpdateReadedMessage rota: /message/{id}/read
func UpdateReadedMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	message, err := repository.UpdateReadedMessage(id)
	responseRequest(w, message, err)
}

// SendMessage Envia mensagens para os funcionários
func SendMessage(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	objMgr := decoderRequest(r, &models.MensagensGenericasReq{})
	Mgr := objMgr.(*models.MensagensGenericasReq)

	uuid, _ := exec.Command("uuidgen").Output()

	var response models.MensagensGenericasRes
	response.IDLote = string(uuid)
	response.Count = len(Mgr.Employees)
	var err error

	var batch models.MessageBatch
	batch.Text = Mgr.Message
	batch.IDUserSend = Mgr.IDUserSend
	batch.IDIntegration = Mgr.IDIntegration

	btc, err := repository.CreateMessageBatch(batch)
	if err != nil {
		responseRequest(w, nil, err)
	}

	for _, mg := range Mgr.Employees {
		var send models.Send
		var error models.Errors

		user, err := buscaWorkplaceUser(mg.Email, Mgr.IDIntegration)

		if err != nil {
			error.EmployeeID = ""
			error.Message = "Erro ao recuperar o id workplace pelo email: " + mg.Email
			response.Errors = append(response.Errors, error)
		} else {
			m, err := sendTextMessage(user.ID, Mgr.Message, Mgr.IDIntegration)

			if err != nil {
				error.EmployeeID = user.ID
				error.Message = "Erro ao encaminhar a mensagem para o colaborador. "
				response.Errors = append(response.Errors, error)
			} else {
				var message models.Message
				message.IDBatch = btc
				message.IDWorkplace = user.ID
				repository.CreateMessage(message)

				send.EmployeeID = user.ID
				send.MessageID = m.MessageID
				send.RecipientID = m.RecipientID
				response.Send = append(response.Send, send)
			}
		}
	}

	responseRequest(w, response, err)
}
