package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestHealth(t *testing.T) {

	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)
}

func TestCreateRule(t *testing.T) {
	table := "rules"
	clearTable(table)

	payload := []byte(`{"typeRules":"action", "description":"Ao logar"}`)

	request, _ := http.NewRequest("POST", "/rules", bytes.NewBuffer(payload))
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)

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
	clearTable(table)
}

func TestGetAllRules(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/rules", nil)
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdateRules(t *testing.T) {
	table := "rules"
	addRules(1)
	payload := []byte(`{"idRules":1, "typeRules": "schedule", "description": " " }`)

	request, _ := http.NewRequest(http.MethodPut, "/rules", bytes.NewBuffer(payload))
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)
	clearTable(table)
}

func TestGetRuleByAny(t *testing.T) {
	table := "rules"
	addRules(1)

	request, _ := http.NewRequest(http.MethodGet, "/rules/1", nil)
	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)
	clearTable(table)
}
