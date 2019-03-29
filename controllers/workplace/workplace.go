package workplace

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"bitbucket.org/magazine-ondemand/adminplace-api/repository"

	"bitbucket.org/magazine-ondemand/adminplace-api/config"
	"bitbucket.org/magazine-ondemand/adminplace-api/models"
)

//BuscaWorkplaceUser é uma função para  busca de um email no workplace
func BuscaWorkplaceUser(email string, idIntegration int) (*models.WPUser, error) {
	var u models.WPUser
	config := config.Configuracoes()
	url := fmt.Sprintf(config.GraphURL + email + "?fields=first_name")

	integration, err := repository.GetIntegrationByID(idIntegration)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+integration.Token)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	if strconv.Itoa(resp.StatusCode) == "200" {
		jsonErr := json.Unmarshal(data, &u)
		if jsonErr != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	return &u, nil
}

//SendTextMessage envia a messagem para um colaborador no workplace
func SendTextMessage(id string, text string, idIntegration int) (*models.MessageResponse, error) {
	var mr models.MessageSend
	var ms models.MessageResponse
	config := config.Configuracoes()
	url := fmt.Sprintf(config.GraphURL + "me/messages")

	mr.MessagingType = "RESPONSE"
	mr.Recipient.ID = id
	mr.MessageData.Text = text

	b, _ := json.Marshal(&mr)
	body := bytes.NewBuffer([]byte(b))

	integration, err := repository.GetIntegrationByID(idIntegration)

	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+integration.Token)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	if strconv.Itoa(resp.StatusCode) == "200" {
		jsonErr := json.Unmarshal(data, &ms)
		if jsonErr != nil {
			return nil, err
		}
	} else {
		err := errors.New("Erro ao enviar mensagem")
		return nil, err
	}

	return &ms, nil
}
