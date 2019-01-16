package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"bitbucket.org/dt_souza/adminplace-api/config"
	"bitbucket.org/dt_souza/adminplace-api/models"

	"bitbucket.org/dt_souza/adminplace-api/repository"
	"github.com/gorilla/mux"
)

// GetGroupByID rota: /group/{id}
func GetGroupByID(w http.ResponseWriter, r *http.Request) {
	var g models.Group
	vars := mux.Vars(r)
	idGroup := vars["id"]
	idIntegration, _ := strconv.Atoi(r.URL.Query().Get("idIntegration"))
	integration, err := repository.GetIntegrationByID(idIntegration)
	config := config.Configuracoes()
	url := fmt.Sprintf(config.GraphURL + idGroup)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+integration.Token)
	if err != nil {
		responseRequest(w, g, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		responseRequest(w, g, err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	if strconv.Itoa(resp.StatusCode) == "200" {
		jsonErr := json.Unmarshal(data, &g)
		if jsonErr != nil {
			responseRequest(w, g, err)
		}
	} else {
		err = errors.New("Erro ao consultar grupo por ID")
		responseRequest(w, g, err)
	}

	responseRequest(w, g, err)
}

// GetAllGroup rota: /group
func GetAllGroup(w http.ResponseWriter, r *http.Request) {
	var g models.GroupList
	idIntegration, _ := strconv.Atoi(r.URL.Query().Get("idIntegration"))
	integration, err := repository.GetIntegrationByID(idIntegration)
	config := config.Configuracoes()
	url := fmt.Sprintf(config.GraphURL + config.Community + "/groups?limit=5000")

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+integration.Token)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	if strconv.Itoa(resp.StatusCode) == "200" {
		jsonErr := json.Unmarshal(data, &g)
		if jsonErr != nil {
			fmt.Println(jsonErr)
		}
	} else {
		err = errors.New("Erro ao consultar grupos")
		fmt.Println(err)
	}

	responseRequest(w, g, err)
}

//DeleteGroupMembers exclui todos os membros de um grupo com exceção dos administradores
func DeleteGroupMembers(w http.ResponseWriter, r *http.Request) {
	var gm models.GroupMembers
	vars := mux.Vars(r)
	idGroup := vars["id"]
	idIntegration, _ := strconv.Atoi(r.URL.Query().Get("idIntegration"))
	integration, err := repository.GetIntegrationByID(idIntegration)
	var page string

	gm = GetGroupMembers(idGroup, integration.Token, page)
	for _, user := range gm.Data {
		if !user.Administrator {
			DeleteMember(idGroup, user.ID, integration.Token)
		}
	}
	for {
		if gm.Paging.Next == "" {
			break
		} else {
			page = gm.Paging.Cursors.After
			gm = GetGroupMembers(idGroup, integration.Token, page)
			for _, user := range gm.Data {
				if !user.Administrator {
					DeleteMember(idGroup, user.ID, integration.Token)
				}
			}
		}
	}

	gm = GetGroupMembers(idGroup, integration.Token, "")
	responseRequest(w, gm, err)
}

func GetGroupMembers(idGroup string, token string, page string) models.GroupMembers {
	var gm models.GroupMembers
	config := config.Configuracoes()
	url := fmt.Sprintf(config.GraphURL + idGroup + "/members?after=" + page)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	if strconv.Itoa(resp.StatusCode) == "200" {
		jsonErr := json.Unmarshal(data, &gm)
		if jsonErr != nil {
			fmt.Println(jsonErr)
		}
	} else {
		err = errors.New("Erro ao consultar grupo por ID")
		fmt.Println(err)
	}

	return gm
}

func DeleteMember(idGroup string, idUser string, token string) {
	config := config.Configuracoes()
	url := fmt.Sprintf(config.GraphURL + idGroup + "/members/" + idUser)

	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}
