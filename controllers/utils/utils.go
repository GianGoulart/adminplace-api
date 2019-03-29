package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	isNull = "Erro: Objeto vazio ou nulo"
)

//DecoderRequest esta função faz o decode de qualquer struct
func DecoderRequest(request *http.Request, obj interface{}) interface{} {
	body, _ := ioutil.ReadAll(request.Body)

	err := json.Unmarshal(body, &obj)
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()
	return obj
}

//ValidationRequest valida se a request não é nula
func ValidationRequest(w http.ResponseWriter, r *http.Request) {
	if r == nil {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(isNull))
	}
}

//ResponseRequest monta o response para envio
func ResponseRequest(w http.ResponseWriter, obj interface{}, err error) {
	w.Header().Set("content-type", "text/json; charset=utf-8")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(obj)
	}
}
