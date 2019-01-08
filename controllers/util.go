package controllers

import (
	"encoding/json"
	"net/http"
)

const (
	isNull = "Erro: Objeto vazio ou nulo"
)

var decoderRequest = func(request *http.Request, obj interface{}) interface{} {
	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&obj); err != nil {
		panic(err)
	}
	defer request.Body.Close()

	return obj
}

var validationRequest = func(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(isNull))
	}
}

var responseRequest = func(w http.ResponseWriter, obj interface{}, err error) {
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(obj)
	}
}
