package rules

import (
	"net/http"

	"bitbucket.org/magazine-ondemand/adminplace-api/repository"

	"bitbucket.org/magazine-ondemand/adminplace-api/models"

	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/utils"
)

//GetAllRules retorna todas as regras cadastradas
func GetAllRules(w http.ResponseWriter, r *http.Request) {
	integration, err := repository.GetAllRules()
	utils.ResponseRequest(w, integration, err)
}

//GetRuleByAny faz busca de regra por qualquer parametro
func GetRuleByAny(w http.ResponseWriter, r *http.Request) {

}

//CreateRule cria uma nova regra
func CreateRule(w http.ResponseWriter, r *http.Request) {
	utils.ValidationRequest(w, r)
	obj := utils.DecoderRequest(r, &models.Rules{})
	b := obj.(*models.Rules)

	rules, err := repository.CreateRules(b)
	b.IDRules = rules
	utils.ResponseRequest(w, b, err)
}

//UpdateRule altera uma regra existente
func UpdateRule(w http.ResponseWriter, r *http.Request) {
	utils.ValidationRequest(w, r)
	obj := utils.DecoderRequest(r, &models.Rules{})

	b := obj.(*models.Rules)

	rules, err := repository.UpdateRules(b)
	utils.ResponseRequest(w, rules, err)
}

//GetAllCrons retorna todos os agendamentos cadastrados
func GetAllCrons(w http.ResponseWriter, r *http.Request) {
	rules, err := repository.GetAllCrons()
	utils.ResponseRequest(w, rules, err)
}

//CreateCron cria um novo agendamento
func CreateCron(w http.ResponseWriter, r *http.Request) {
	utils.ValidationRequest(w, r)
	obj := utils.DecoderRequest(r, &models.Cron{})
	b := obj.(*models.Cron)

	cron, err := repository.CreateCrons(b)
	utils.ResponseRequest(w, cron, err)
}
