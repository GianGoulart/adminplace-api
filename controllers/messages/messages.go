package messages

import (
	"fmt"
	"net/http"
	"strconv"

	"bitbucket.org/magazine-ondemand/adminplace-api/models"
	"bitbucket.org/magazine-ondemand/adminplace-api/repository"

	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/groups"
	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/utils"
	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/workplace"

	"github.com/gorilla/mux"
)

// GetMessageByID rota: /message/{id}
func GetMessageByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	message, err := repository.GetMessageByID(id)
	utils.ResponseRequest(w, message, err)
}

// GetMessageByUser rota: /message/{user}
func GetMessageByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, _ := strconv.Atoi(vars["user"])
	message, err := repository.GetMessageByUser(user)
	utils.ResponseRequest(w, message, err)
}

// CreateMessage rota: /message
func CreateMessage(w http.ResponseWriter, r *http.Request) {
	utils.ValidationRequest(w, r)
	obj := utils.DecoderRequest(r, &models.Message{})
	msg := obj.(models.Message)

	message, err := repository.CreateMessage(msg)
	utils.ResponseRequest(w, message, err)
}

/*  UpdateReceivedMessage rota: /message/{id}/receive
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
*/

// SendMessage Envia mensagens para os funcionários
func SendMessage(w http.ResponseWriter, r *http.Request) {
	utils.ValidationRequest(w, r)
	objMgr := utils.DecoderRequest(r, &models.MensagensGenericasReq{})
	Mgr := objMgr.(*models.MensagensGenericasReq)

	var response models.MensagensGenericasRes

	response.Count = len(Mgr.Employees)
	var err error

	var batch models.MessageBatch
	batch.Text = Mgr.Message
	batch.IDUserSend = Mgr.IDUserSend
	batch.IDIntegration = Mgr.IDIntegration

	btc, err := repository.CreateMessageBatch(batch)
	if err != nil {
		utils.ResponseRequest(w, nil, err)
	}

	response.IDLote = string(btc)

	for _, mg := range Mgr.Employees {
		var send models.Send
		var error models.Errors

		user, err := workplace.BuscaWorkplaceUser(mg.Email, Mgr.IDIntegration)

		if err != nil {
			error.EmployeeID = ""
			error.Message = "Erro ao recuperar o id workplace pelo email: " + mg.Email
			response.Errors = append(response.Errors, error)
		} else {
			m, err := workplace.SendTextMessage(user.ID, Mgr.Message, Mgr.IDIntegration)

			if err != nil {
				error.EmployeeID = user.ID
				error.Message = "Erro ao encaminhar a mensagem para o colaborador. "
				response.Errors = append(response.Errors, error)
			} else {
				var message models.Message
				message.IDBatch = btc
				message.IDWorkplace = user.ID
				fmt.Println(message)
				repository.CreateMessage(message)

				send.EmployeeID = user.ID
				send.MessageID = m.MessageID
				send.RecipientID = m.RecipientID
				response.Send = append(response.Send, send)
			}
		}
	}

	utils.ResponseRequest(w, response, err)
}

// SendGroupMessage Envia mensagens para os funcionários de um grupo
func SendGroupMessage(w http.ResponseWriter, r *http.Request) {
	utils.ValidationRequest(w, r)
	objGm := utils.DecoderRequest(r, &models.GroupMessage{})
	gm := objGm.(*models.GroupMessage)

	var response models.MensagensGenericasRes
	var err error

	var batch models.MessageBatch
	batch.Text = gm.Text
	batch.IDUserSend = gm.IDUserSend
	batch.IDIntegration = gm.IDIntegration

	btc, err := repository.CreateMessageBatch(batch)
	if err != nil {
		utils.ResponseRequest(w, nil, err)
	}
	response.IDLote = string(btc)
	integration, err := repository.GetIntegrationByID(gm.IDIntegration)
	var gms models.GroupMembers

	var page string
	var send models.Send
	var error models.Errors

	gms = groups.GetGroupMembers(gm.IDGroup, integration.Token, page)
	for _, user := range gms.Data {
		m, err := workplace.SendTextMessage(user.ID, gm.Text, gm.IDIntegration)

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
	for {
		if gms.Paging.Next == "" {
			break
		} else {
			page = gms.Paging.Cursors.After
			gms = groups.GetGroupMembers(gm.IDGroup, integration.Token, page)
			for _, user := range gms.Data {
				m, err := workplace.SendTextMessage(user.ID, gm.Text, gm.IDIntegration)
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
	}
}
