package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"bitbucket.org/dt_souza/adminplace-api/config"
	"bitbucket.org/dt_souza/adminplace-api/models"
)

func buscaWorkplaceUser(email string) (*models.WPUser, error) {
	var u models.WPUser
	config := config.Configuracoes()
	url := fmt.Sprintf(config.GraphURL + email + "?fields=first_name")

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", config.PageAccessToken)
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

func sendTextMessage(id string, text string) (*models.MessageResponse, error) {
	var mr models.MessageSend
	var ms models.MessageResponse
	config := config.Configuracoes()
	url := fmt.Sprintf(config.GraphURL + "me/messages")

	mr.MessagingType = "RESPONSE"
	mr.Recipient.ID = id
	mr.MessageData.Text = text

	b, _ := json.Marshal(&mr)
	body := bytes.NewBuffer([]byte(b))

	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", config.PageAccessToken)
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
		return nil, err
	}

	return &ms, nil
}
