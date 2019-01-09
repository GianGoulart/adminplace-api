package controllers

import (
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
	var err error

	for _, mg := range Mgr.Employees {
		var send models.Send
		var error models.Errors

		user, err := buscaWorkplaceUser(mg.Email)

		if err != nil {
			error.EmployeeID = ""
			error.Message = "Erro ao recuperar o id workplace pelo email: " + mg.Email
			response.Errors = append(response.Errors, error)
		} else {
			m, err := sendTextMessage(user.ID, Mgr.Message)

			if err != nil {
				error.EmployeeID = user.ID
				error.Message = "Erro ao encaminhar a mensagem para o colaborador. "
				response.Errors = append(response.Errors, error)
			} else {
				send.EmployeeID = user.ID
				send.MessageID = m.MessageID
				send.RecipientID = m.RecipientID
				response.Send = append(response.Send, send)
			}
		}
	}

	responseRequest(w, response, err)
}
