package controllers

import (
	"net/http"

	"bitbucket.org/magazine-ondemand/adminplace-api/repository"

	"bitbucket.org/magazine-ondemand/adminplace-api/models"
)

//GetAllRules retorna todas as regras cadastradas
func GetAllRules(w http.ResponseWriter, r *http.Request) {
	integration, err := repository.GetAllRules()
	responseRequest(w, integration, err)
}

//GetRuleByAny faz busca de regra por qualquer parametro
func GetRuleByAny(w http.ResponseWriter, r *http.Request) {

}

//CreateRule cria uma nova regra
func CreateRule(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	obj := decoderRequest(r, &models.Rules{})
	b := obj.(*models.Rules)

	rules, err := repository.CreateRules(b)
	b.IDRules = rules
	responseRequest(w, b, err)
}

//UpdateRule altera uma regra existente
func UpdateRule(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	obj := decoderRequest(r, &models.Rules{})

	b := obj.(*models.Rules)

	rules, err := repository.UpdateRules(b)
	responseRequest(w, rules, err)
}

//GetAllCrons retorna todos os agendamentos cadastrados
func GetAllCrons(w http.ResponseWriter, r *http.Request) {
	rules, err := repository.GetAllCrons()
	responseRequest(w, rules, err)
}

//CreateCron cria um novo agendamento
func CreateCron(w http.ResponseWriter, r *http.Request) {
	validationRequest(w, r)
	obj := decoderRequest(r, &models.Cron{})
	b := obj.(*models.Cron)

	cron, err := repository.CreateCrons(b)
	responseRequest(w, cron, err)
}
