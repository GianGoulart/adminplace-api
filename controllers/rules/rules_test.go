package rules_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreateRule(t *testing.T) {
	table := "rules"
	ClearTable(table)

	payload := []byte(`{"typeRules":"action", "description":"Ao logar"}`)

	request, _ := http.NewRequest("POST", "/rules", bytes.NewBuffer(payload))
	response := ExecuteRequest(request)

	CheckResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["typeRules"] != "action" {
		t.Errorf("Expected type rules to be 'action'. Got '%v'", m["typeRules"])
	}

	if m["description"] != "Ao logar" {
		t.Errorf("Expected description to be 'Ao logar'. Got '%v'", m["description"])
	}

	if m["idRules"] != 1.0 {
		t.Errorf("Expected product idRules to be '1'. Got '%v'", m["idRules"])
	}
	ClearTable(table)
}

func TestGetAllRules(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/rules", nil)
	response := ExecuteRequest(request)

	CheckResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateRules(t *testing.T) {
	table := "rules"
	AddRules(1)
	payload := []byte(`{"idRules":1, "typeRules": "schedule", "description": " " }`)

	request, _ := http.NewRequest(http.MethodPut, "/rules", bytes.NewBuffer(payload))
	response := ExecuteRequest(request)

	CheckResponseCode(t, http.StatusOK, response.Code)
	ClearTable(table)
}

func TestGetRuleByAny(t *testing.T) {
	table := "rules"
	AddRules(1)

	request, _ := http.NewRequest(http.MethodGet, "/rules/1", nil)
	response := ExecuteRequest(request)

	CheckResponseCode(t, http.StatusOK, response.Code)
	ClearTable(table)
}

func TestGetAllCrons(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/crons", nil)
	response := ExecuteRequest(request)

	CheckResponseCode(t, http.StatusOK, response.Code)

}

func TestCreateCron(t *testing.T) {

}
