package controllers

import (
	"encoding/json"
	"net/http"
)

//HealthCheck ...
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Gerando um objeto customizado à partir de um map, e o convertendo em json
	response, _ := json.Marshal(map[string]interface{}{
		"status": "up"})
	// Write escreve o conteúdo do slice de bytes no corpo da resposta
	w.WriteHeader(http.StatusOK)
	w.Write(response)

	return
}
