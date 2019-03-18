package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	isNull = "Erro: Objeto vazio ou nulo"
)

var decoderRequest = func(request *http.Request, obj interface{}) interface{} {
	body, _ := ioutil.ReadAll(request.Body)

	err := json.Unmarshal(body, &obj)
	if err != nil {
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
	w.Header().Set("content-type", "text/json; charset=utf-8")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(obj)
	}
}
