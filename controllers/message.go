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
		var envio models.Envios
		var erro models.Erros

		user, err := buscaWorkplaceUser(mg.Email)

		if err != nil {
			erro.EmployeeID = user.ID
			erro.Message = "Erro ao recuperar o id workplace do colaborador. " + err.DeveloperMessage
			response.Erros = append(response.Erros, erro)
		}

		if user.ID != "" {
			ms, err := sendTextMessage(user.ID, Mgr.Message)

			if err != nil {
				erro.EmployeeID = user.ID
				erro.Message = "Erro ao encaminhar a mensagem para o colaborador. " + err.DeveloperMessage
				response.Erros = append(response.Erros, erro)
			} else {
				envio.EmployeeID = user.ID
				envio.MessageID = ms.MessageID
				envio.RecipientID = ms.RecipientID
				response.Envios = append(response.Envios, envio)
			}
		} else {
			erro.EmployeeID = user.ID
			erro.Message = "Id workplace do colaborador não encontrado."
			response.Erros = append(response.Erros, erro)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
