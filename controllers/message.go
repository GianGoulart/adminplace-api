package controllers

import (
	"encoding/json"
	"net/http"
	"os/exec"

	"bitbucket.org/dt_souza/adminplace-api/models"
)

// SendMessage Envia mensagens para os funcionários
func SendMessage(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	objMgr := decoderRequest(r, &models.MensagensGenericasReq{})
	Mgr := objMgr.(*models.MensagensGenericasReq)

	uuid, _ := exec.Command("uuidgen").Output()

	var response models.MensagensGenericasRes
	response.IDLote = string(uuid)
	response.Count = len(Mgr.Employees)

	for _, mg := range Mgr.Employees {
		var send models.Send
		var error models.Errors

		user, err := buscaWorkplaceUser(mg.Email)

		if err != nil {
			error.EmployeeID = ""
			error.Message = "Erro ao recuperar o id workplace pelo email: " + mg.Email
			response.Errors = append(response.Errors, error)
		} else {
			ms, err := sendTextMessage(user.ID, Mgr.Message)

			if err != nil {
				error.EmployeeID = user.ID
				error.Message = "Erro ao encaminhar a mensagem para o colaborador. " + err.DeveloperMessage
				response.Errors = append(response.Errors, error)
			} else {
				send.EmployeeID = user.ID
				send.MessageID = ms.MessageID
				send.RecipientID = ms.RecipientID
				response.Send = append(response.Send, send)
			}
		}
	}

	w.Header().Set("content-type", "text/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
